package main

import (
	"context"
	"fmt"

	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error: cannot get user data from database: %v", err)
		}

		return handler(s, cmd, currentUser)
	}
}
