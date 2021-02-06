package repository

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
)

type ActionRepository interface {
	Repository
	UpdateActionStatus(ctx context.Context, actionID uint, status models.ActionStatus) error
	FetchActions(ctx context.Context) ([]models.Action, error)
}
