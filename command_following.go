package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	feed, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error: cannot get feed data from database: %v", err)
	}

	fmt.Printf("Posts you follow:\n")
	for _, post := range feed {
		fmt.Printf("%s\n", post.FeedName)
	}

	return nil
}
