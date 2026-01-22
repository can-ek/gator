package main

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/can-ek/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user db.User) error {
	var limit int32 = 2

	if len(cmd.args) == 1 {
		param, err := strconv.Atoi(cmd.args[0])

		// If there's an error while parsing, keep default
		if err == nil {
			limit = int32(param)
		}
	}

	ctx := context.Background()
	params := db.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(ctx, params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println("\t", post.Title, "\t", post.PublishedAt)
	}

	return nil
}
