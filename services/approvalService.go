package services

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
)

type ApprovalServiceProviders interface {
	Update(ctx context.Context, provider string, actionID uint, actionStatus models.ActionStatus) error
}

type ApprovalService interface {
	Update(ctx context.Context, actionID uint, actionStatus models.ActionStatus) error
}
