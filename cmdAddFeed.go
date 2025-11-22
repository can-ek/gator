package main

import (
	"context"
	"fmt"
	"time"

	intdb "github.com/can-ek/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed expects name and url arguments")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.configuration.CurrentUsername)
	if err != nil {
		return err
	}

	currTime := time.Now()
	params := intdb.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, params)
	if err != nil {
		return err
	}

	feed_params := intdb.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, feed_params)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
