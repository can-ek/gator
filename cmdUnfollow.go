package main

import (
	"context"
	"fmt"

	database "github.com/can-ek/gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("unfollow expects url argument")
	}

	ctx := context.Background()
	feed, err := s.db.GetFeedsByUrl(ctx, cmd.args[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowsForUserParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	_, err = s.db.DeleteFeedFollowsForUser(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println("User:", user.Name, "Unfollowed:", feed.Name)
	return nil
}
