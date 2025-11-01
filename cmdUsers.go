package main

import (
	"context"
	"fmt"
)

func handleUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		name := user.Name
		if user.Name == s.configuration.CurrentUsername {
			name = fmt.Sprintf("%s (current)", name)
		}

		fmt.Println(name)
	}

	return nil
}
