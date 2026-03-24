package main

import (
	"context"
	"fmt"
)

func handlerFetchFeeds(s *state, cmd command) error {
	feeds, err := s.db.FetchFeedsWithUserName(context.Background())
	if err != nil {
		return fmt.Errorf("error while fetching feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("name: %s\nURL: %s\nUser Name: %s\n", feed.Name, feed.Url, feed.UserName.String)
		fmt.Printf("-----------------------------------\n")
	}
	return nil
}
