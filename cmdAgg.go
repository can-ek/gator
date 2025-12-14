package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/can-ek/gator/internal/database"
	rss "github.com/can-ek/gator/rss"
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

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
