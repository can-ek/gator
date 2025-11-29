package main

import (
	"context"
	"fmt"
	"time"

	int_db "github.com/can-ek/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user int_db.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("follow expects url argument")
	}

	url := cmd.args[0]
	ctx := context.Background()
	feed, err := s.db.GetFeedsByUrl(ctx, url)
	if err != nil {
		return err
	}	

	currTime := time.Now()
	params := int_db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println(
		"User:",
		feedFollow.UserName,
		"follows Feed:",
		feedFollow.FeedName)

	return nil
}
