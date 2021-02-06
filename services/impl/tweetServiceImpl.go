package serviceImpl

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"github.com/tiwariayush700/tweeting/services"
	"gorm.io/gorm"
)

type tweetServiceImpl struct {
	repository       repository.TweetRepository
	actionRepository repository.ActionRepository
}

func (t *tweetServiceImpl) Update(ctx context.Context, actionID uint, actionStatus models.ActionStatus) error {

	action := &models.Action{}
	err := t.actionRepository.Get(ctx, action, actionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userError.ErrorNotFound
		}
		return err
	}

	tweet := &models.Tweet{}
	err = json.Unmarshal(action.Body, tweet)
	if err != nil {
		logrus.Errorf("err unmarshalling action body err : %v", err)
		return err
	}

	err = t.repository.UpdateTweetMessage(ctx, tweet.ID, tweet.Message)
	if err != nil {
		logrus.Errorf("err updating user role err : %v", err)
		return err
	}

	return t.actionRepository.UpdateActionStatus(ctx, actionID, actionStatus)
}

func (t *tweetServiceImpl) GetTweetsByUserID(ctx context.Context, userID uint) ([]models.Tweet, error) {
	return t.repository.GetTweetsByUserID(ctx, userID)
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

func NewTweetApprovalServiceImpl(repository repository.TweetRepository, actionRepository repository.ActionRepository) services.ApprovalService {
	return &tweetServiceImpl{repository: repository, actionRepository: actionRepository}
}
