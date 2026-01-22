package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/can-ek/gator/internal/database"
	rss "github.com/can-ek/gator/rss"
	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("agg expects time_between_reqs argument")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	currTime := time.Now()
	params := db.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  currTime,
			Valid: true,
		},
		ID: feed.ID,
	}

	feed, err = s.db.MarkFeedFetched(ctx, params)
	if err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	fmt.Println("Found", len(rssFeed.Channel.Item), "posts")

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{String: "", Valid: false}

		if item.Description != "" {
			desc = sql.NullString{String: item.Description, Valid: true}
		}

		parsedDate, _ := time.Parse(time.RFC1123Z, item.PubDate)

		postParams := db.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currTime,
			UpdatedAt:   currTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: parsedDate,
			FeedID:      feed.ID,
		}

		_, err := s.db.CreatePost(ctx, postParams)
		if err != nil {
			return err
		}
	}

	return nil
}
