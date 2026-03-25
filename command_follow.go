package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("the follow handler expect one argument, the url")
	}

	feed, err := s.db.GetFeedsByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error: cannot get feed data from database: %v", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	fmt.Printf("user: %s is now following %s\n", follow.UserName, follow.FeedName)
	return nil
}
