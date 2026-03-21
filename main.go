package main

import (
	"log"
	"os"

	"github.com/viniciusLambert/blog-aggregator/internal/config"
)

func main() {
	cfgData, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error getting config file: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	st := state{
		cfg: &cfgData,
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

	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatalf("failed running command: %v\n", err)
	}
}
