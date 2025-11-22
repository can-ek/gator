package main

import (
	"context"
	"fmt"
)

func handleFollowing(s *state, cmd command) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.configuration.CurrentUsername)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println("User:", user.Name, "following:")
	for _, feed := range feeds {
		fmt.Println("\t", feed.Name)
	}

	return nil
}
