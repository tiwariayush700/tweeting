package repository

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
)

type TweetRepository interface {
	Repository
	FetchTweets(ctx context.Context) ([]models.Tweet, error)
	UpdateTweetMessage(ctx context.Context, tweetID uint, message string) error
}
