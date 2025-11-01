package main

import "context"

func handleReset(s *state, cmd command) error {
	return s.db.ResetUsers(context.Background())
}
