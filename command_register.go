package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the register handles expects a single argument, the user name")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("error inserting user on database: %v", err)
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("failed to set user. %v", err)
	}

	if err := s.cfg.WriteFile(); err != nil {
		return fmt.Errorf("failed to write file. %v", err)
	}

	fmt.Printf("user %s was created\nUser data: %v\n", cmd.Args[0], user)

	return nil
}
