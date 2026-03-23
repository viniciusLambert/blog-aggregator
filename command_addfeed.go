package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("the addfeed handler expects two arguments, the feed name and url")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error get user data from database: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.Args[0],
			Url:       cmd.Args[1],
			UserID:    currentUser.ID,
		})
	if err != nil {
		return fmt.Errorf("error creating new feed: %v", err)
	}

	fmt.Println(feed)
	return nil
}
