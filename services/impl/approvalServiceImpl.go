package serviceImpl

import (
	"context"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/services"
)

type approvalServiceImpl struct {
	providers map[string]services.ApprovalService
}

func (a *approvalServiceImpl) Update(ctx context.Context, provider string, actionID uint, actionStatus models.ActionStatus) error {

	if _, ok := a.providers[provider]; !ok {
		panic("provider missing")
	}

	return a.providers[provider].Update(ctx, actionID, actionStatus)
}

func NewApprovalServiceImpl(providers map[string]services.ApprovalService) services.ApprovalServiceProviders {
	return &approvalServiceImpl{providers: providers}
}
