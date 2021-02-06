package repositoryImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"gorm.io/gorm"
)

type tweetRepositoryImpl struct {
	repositoryImpl
}

func (t *tweetRepositoryImpl) FetchTweets(ctx context.Context) ([]models.Tweet, error) {

	tweets := make([]models.Tweet, 0)

	err := t.DB.Find(&tweets).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}
		return nil, err
	}

	return tweets, nil
}

func (t *tweetRepositoryImpl) UpdateTweetMessage(ctx context.Context, tweetID uint, message string) error {

	err := t.DB.Model(&models.Action{}).
		Where("id = ?", tweetID).
		Update("message", message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userError.ErrorNotFound
		}
		return err
	}

	return nil
}

func NewTweetRepositoryImpl(db *gorm.DB) repository.TweetRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &tweetRepositoryImpl{repoImpl}
}
