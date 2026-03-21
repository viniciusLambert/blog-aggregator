package main

import (
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the login handles expects a single argument, the username")
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		log.Fatalf("failed to set user. %v\n", err)
	}

	if err := s.cfg.WriteFile(); err != nil {
		log.Fatalf("failed to write file. %v\n", err)
	}

	fmt.Printf("Username has been seted to %s\n", cmd.Args[0])

	return nil
}
