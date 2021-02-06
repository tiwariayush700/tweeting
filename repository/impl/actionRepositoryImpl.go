package repositoryImpl

import (
	"github.com/tiwariayush700/tweeting/repository"
	"gorm.io/gorm"
)

type actionRepositoryImpl struct {
	repositoryImpl //overrides basic CRUD repo
}

func NewActionRepositoryImpl(db *gorm.DB) repository.ActionRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &actionRepositoryImpl{repoImpl}
}
