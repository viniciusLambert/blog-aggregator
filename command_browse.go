package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/viniciusLambert/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit string
	if len(cmd.Args) > 0 {
		limit = cmd.Args[0]
	} else {
		limit = "2"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("invalid limit: %v", err)
	}

	posts, err := s.db.GetPostForUser(context.Background(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  int32(limitInt),
	})
	if err != nil {
		return fmt.Errorf("error while getting posts: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("%s\nPubDate: %v\n%v\n-------------------------\n", post.Title, post.PublishedAt, post.Description)
	}

	return nil
}
