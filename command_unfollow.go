package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("the follow handler expect one argument, the url")
	}

	feed, err := s.db.GetFeedsByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error: cannot get feed data from database: %v", err)
	}

	if err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("error while deleting follow status from database: %v", err)
	}

	return nil
}
