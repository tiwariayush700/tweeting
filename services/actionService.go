package services

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
)

type ActionService interface {
	CreateAction(ctx context.Context, action *models.Action) error
	GetActionByID(ctx context.Context, actionID uint) (*models.Action, error)
	FetchActions(ctx context.Context) ([]models.Action, error)
}
