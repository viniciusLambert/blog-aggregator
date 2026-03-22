package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the login handles expects a single argument, the username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to retrieve user from database: %v", err)
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("failed to set user. %v", err)
	}

	if err := s.cfg.WriteFile(); err != nil {
		return fmt.Errorf("failed to write file. %v", err)
	}

	fmt.Printf("Username has been seted to %s\n", cmd.Args[0])

	return nil
}
