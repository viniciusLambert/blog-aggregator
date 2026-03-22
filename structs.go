package main

import (
	"fmt"

	"github.com/viniciusLambert/blog-aggregator/internal/config"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	RegisteredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, exist := c.RegisteredCommands[cmd.Name]

	if !exist {
		return fmt.Errorf("function does not exist: %v", cmd.Name)
	}
	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.RegisteredCommands[name] = f
}
