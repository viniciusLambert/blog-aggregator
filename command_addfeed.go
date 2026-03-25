package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("the addfeed handler expects two arguments, the feed name and url")
	}

	feed, err := s.db.CreateFeed(context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.Args[0],
			Url:       cmd.Args[1],
			UserID:    user.ID,
		})
	if err != nil {
		return fmt.Errorf("error creating new feed: %v", err)
	}

	cmd.Args[0] = cmd.Args[1]
	if err = handlerFollow(s, cmd, user); err != nil {
		return fmt.Errorf("error while creating follow bond: %v", err)
	}

	fmt.Println(feed)
	return nil
}
