package main

import "fmt"

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login expects username argument")
	}

	username := cmd.args[0]
	err := s.configuration.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User set to:", username)
	return nil
}


