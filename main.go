// Package main contains the entry functionality for the program
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	config "github.com/can-ek/gator/config"
	"github.com/can-ek/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	configuration *config.Config
	db            *database.Queries
}

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Println("ERROR: Could not read config:", err)
		return
	}

	db, err := sql.Open("postgres", configuration.DBURL)
	if err != nil {
		fmt.Println("ERROR: Could not connect to Database:", err)
		return
	}

	dbQueries := database.New(db)
	currState := state{configuration: &configuration, db: dbQueries}
	supportedCmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	supportedCmds.register("login", handleLogin)
	supportedCmds.register("register", handleRegister)
	supportedCmds.register("reset", handleReset)
	supportedCmds.register("users", handleUsers)
	supportedCmds.register("agg", handleAgg)
	supportedCmds.register("addfeed", handleAddFeed)
	supportedCmds.register("feeds", handleFeeds)
	supportedCmds.register("follow", handleFollow)
	supportedCmds.register("following", handleFollowing)

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
