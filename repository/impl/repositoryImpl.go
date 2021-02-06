package repositoryImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	DB *gorm.DB
}

func (r *repositoryImpl) Create(ctx context.Context, out interface{}) error {
	err := r.DB.Create(out).Error
	if err == gorm.ErrRecordNotFound {
		return userError.ErrorNotFound
	}

	return err
}

func (r *repositoryImpl) Get(ctx context.Context, out interface{}, id interface{}) error {
	err := r.DB.First(out, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return userError.ErrorNotFound
	}

	return err
}

func NewRepositoryImpl(db *gorm.DB) repository.Repository {
	return &repositoryImpl{DB: db}
}
