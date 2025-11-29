package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.cmds[cmd.name]; exists {
		return handler(s, cmd)
	} else {
		return fmt.Errorf("ERROR: Command %s not found", cmd.name)
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
