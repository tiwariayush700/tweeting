package serviceImpl

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/repository"
	"github.com/tiwariayush700/tweeting/services"
	"github.com/tiwariayush700/tweeting/utils"
	"gorm.io/gorm"
)

type userServiceImpl struct {
	repository       repository.UserRepository
	actionRepository repository.ActionRepository
}

func (u *userServiceImpl) Update(ctx context.Context, actionID uint, actionStatus models.ActionStatus) error {

	action := &models.Action{}
	err := u.actionRepository.Get(ctx, action, actionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userError.ErrorNotFound
		}
		return err
	}

	user := &models.User{}
	err = json.Unmarshal(action.Body, user)
	if err != nil {
		logrus.Errorf("err unmarshalling action body err : %v", err)
		return err
	}

	err = u.UpdateUserRole(ctx, user.ID, user.Role)
	if err != nil {
		logrus.Errorf("err updating user role err : %v", err)
		return err
	}

	return u.actionRepository.UpdateActionStatus(ctx, actionID, actionStatus)
}

func (u *userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	user.Password = utils.GetMd5(user.Password)

	err := u.repository.Create(ctx, user)
	return err
}

func (u *userServiceImpl) GetUserByID(ctx context.Context, userID uint) (*models.UserResponse, error) {

	user := &models.User{}
	err := u.repository.Get(ctx, user, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}

		return nil, err
	}

	userResponse, err := mapUserResponse(user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *userServiceImpl) LoginUser(ctx context.Context, email, password string) (*models.UserResponse, error) {
	password = utils.GetMd5(password)

	user, err := u.repository.GetUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}

	userResponse, err := mapUserResponse(user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *userServiceImpl) UpdateUserRole(ctx context.Context, userId uint, role string) error {

	return u.repository.UpdateUserRole(ctx, userId, role)

}

func NewUserServiceImpl(repository repository.UserRepository) services.UserService {
	return &userServiceImpl{repository: repository}
}

func mapUserResponse(user *models.User) (*models.UserResponse, error) {

	userBytes, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("err marshalling user : err %v", err)
		return nil, err
	}

	userResponse := &models.UserResponse{}
	err = json.Unmarshal(userBytes, userResponse)
	if err != nil {
		logrus.Errorf("err unmarshalling user : err %v", err)
		return nil, err
	}

	return userResponse, nil
}

func NewUserApprovalServiceImpl(repository repository.UserRepository, actionRepository repository.ActionRepository) services.ApprovalService {
	return &userServiceImpl{repository: repository, actionRepository: actionRepository}
}
