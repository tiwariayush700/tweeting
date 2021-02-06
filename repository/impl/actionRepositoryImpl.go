package repositoryImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"gorm.io/gorm"
)

type actionRepositoryImpl struct {
	repositoryImpl //overrides basic CRUD repo
}

func (a *actionRepositoryImpl) FetchActions(ctx context.Context) ([]models.Action, error) {

	actions := make([]models.Action, 0)

	err := a.DB.Find(&actions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}
		return nil, err
	}

	return actions, nil
}

func (a *actionRepositoryImpl) UpdateActionStatus(ctx context.Context, actionID uint, status models.ActionStatus) error {

	err := a.DB.Model(&models.Action{}).
		Where("id = ?", actionID).
		Update("status", status).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userError.ErrorNotFound
		}
		return err
	}

	return nil
}

func NewActionRepositoryImpl(db *gorm.DB) repository.ActionRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &actionRepositoryImpl{repoImpl}
}
