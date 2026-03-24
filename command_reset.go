package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error while reseting database: %v", err)
	}

	fmt.Printf("Users table was reseted\n")

	err = s.db.ClearFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error while reseting database: %v", err)
	}

	fmt.Printf("Feeds table was reseted\n")
	return nil
}
