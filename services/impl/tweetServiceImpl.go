package serviceImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"github.com/tiwariayush700/tweeting/services"
	"gorm.io/gorm"
)

type tweetServiceImpl struct {
	repository repository.TweetRepository
}

func (t *tweetServiceImpl) FetchTweets(ctx context.Context) ([]models.Tweet, error) {
	return t.repository.FetchTweets(ctx)
}

func (t *tweetServiceImpl) CreateTweet(ctx context.Context, tweet *models.Tweet) error {
	return t.repository.Create(ctx, tweet)
}

func (t *tweetServiceImpl) GetTweetByID(ctx context.Context, tweetID uint) (*models.Tweet, error) {

	tweet := &models.Tweet{}
	err := t.repository.Get(ctx, tweet, tweetID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}

		return nil, err
	}

	return tweet, nil
}

func NewTweetServiceImpl(repository repository.TweetRepository) services.TweetService {
	return &tweetServiceImpl{repository: repository}
}
