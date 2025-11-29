package main

import (
	"context"
	"fmt"
	database "github.com/can-ek/gator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()
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
