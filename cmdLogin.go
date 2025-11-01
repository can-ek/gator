package main

import (
	"context"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects username argument")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.configuration.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User set to:", username)
	return nil
}
