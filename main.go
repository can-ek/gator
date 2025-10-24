// Package main contains the entry functionality for the program
package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/can-ek/gator/config"
)

type state struct {
	configuration *config.Config
}

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Println("ERROR: Could not read config:", err)
		return
	}

	currState := state{configuration: &configuration}
	supportedCmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	supportedCmds.register("login", handleLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("ERROR: Missing command name")
		return
	}

	cmd := command{name: args[1], args: args[2:]}
	err = supportedCmds.run(&currState, cmd)
	if err != nil {
		msg := fmt.Sprintf(
			"ERROR: While running %s command, %v",
			cmd.name,
			err)

		log.Fatal(msg)
	}
}
