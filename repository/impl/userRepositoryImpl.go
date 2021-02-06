package repositoryImpl

import (
	"context"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	repositoryImpl //overrides basic CRUD repo
}

func (u *userRepositoryImpl) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*models.User, error) {

	user := &models.User{}
	err := u.DB.Where("email = ? AND password = ?", email, password).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) UpdateUserRole(ctx context.Context, userId uint, role string) error {

	err := u.DB.Model(&models.User{}).
		Where("id = ?", userId).
		Update("role", role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userError.ErrorNotFound
		}
		return err
	}

	return nil
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &userRepositoryImpl{repoImpl}
}
