package serviceImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"github.com/tiwariayush700/tweeting/services"
	"gorm.io/gorm"
)

type actionServiceImpl struct {
	repository repository.ActionRepository
}

func (a *actionServiceImpl) CreateAction(ctx context.Context, action *models.Action) error {
	return a.repository.Create(ctx, action)
}

func (a *actionServiceImpl) GetActionByID(ctx context.Context, actionID uint) (*models.Action, error) {

	action := &models.Action{}
	err := a.repository.Get(ctx, action, actionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}

		return nil, err
	}

	return action, nil
}

func NewActionServiceImpl(repository repository.ActionRepository) services.ActionService {
	return &actionServiceImpl{repository: repository}
}
