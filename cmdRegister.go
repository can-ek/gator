package main

import (
	"context"
	"fmt"
	"time"

	intdb "github.com/can-ek/gator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register expects name argument")
	}

	name := cmd.args[0]
	ctx := context.Background()
	user, _ := s.db.GetUser(ctx, name)
	if user != (intdb.User{}) {
		return fmt.Errorf("user already exists")
	}

	currTime := time.Now()
	params := intdb.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      name,
	}

	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	s.configuration.SetUser(user.Name)
	fmt.Println("User registered successfully")
	return nil
}
