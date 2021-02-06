package repository

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, out interface{}) error
	Get(ctx context.Context, out interface{}, id interface{}) error
	//update
	//delete
}
