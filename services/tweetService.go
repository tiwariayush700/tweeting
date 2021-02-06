package services

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *models.Tweet) error
	GetTweetByID(ctx context.Context, tweetID uint) (*models.Tweet, error)
	FetchTweets(ctx context.Context) ([]models.Tweet, error)
	GetTweetsByUserID(ctx context.Context, userID uint) ([]models.Tweet, error)
}
