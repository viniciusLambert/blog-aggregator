package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/viniciusLambert/blog-aggregator/internal/config"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func main() {
	st, err := SetupState()
	if err != nil {
		log.Fatalf("error while setup state: %v\n", err)
	}

	args := os.Args
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	cmds := commands{
		RegisteredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatalf("failed running command: %v\n", err)
	}
}

func SetupState() (state, error) {
	cfgData, err := config.ReadConfig()
	if err != nil {
		return state{}, fmt.Errorf("error getting config file: %v", err)
	}

	db, err := sql.Open("postgres", cfgData.DbURL)
	if err != nil {
		return state{}, fmt.Errorf("error connecting database: %v", err)
	}
	dbQueries := database.New(db)

	if len(os.Args) < 2 {
		return state{}, fmt.Errorf("not enough arguments")
	}
	return state{
		db:  dbQueries,
		cfg: &cfgData,
	}, nil
}
