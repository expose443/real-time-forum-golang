package service

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/models"
	"github.com/expose443/real-time-forum-golang/backend-api/internal/repository"
)

type AuthService interface {
	IsValidUser(identifier any, password string) (models.User, error)
	CreateUser(user *models.User) error
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		userRepo: dao.NewUserRepo(),
	}
}

type authService struct {
	userRepo repository.UserRepository
}

func (a *authService) IsValidUser(identifier any, password string) (models.User, error) {
	t := reflect.TypeOf(identifier)

	switch t {
	case reflect.TypeOf(""):
		user, err := a.userRepo.GetUserByIdentifier(fmt.Sprintf("%v", identifier))
		if err != nil {
			return models.User{}, err
		}
		if user.Password != password {
			return models.User{}, errors.New("invalid password")
		}
		return user, nil

	case reflect.TypeOf(0):
		num, _ := strconv.Atoi(fmt.Sprintf("%v", identifier))
		user, err := a.userRepo.GetUserByUserID(num)
		if err != nil {
			return models.User{}, err
		}
		if user.Password != password {
			return models.User{}, errors.New("invalid password")
		}
		return user, nil
	default:
		return models.User{}, errors.New("invalid type of identifier")

	}
}

func (a *authService) CreateUser(user *models.User) error {
	_, err := a.userRepo.GetUserByIdentifier(user.Email)
	if err == nil {
		return errors.New("user exists")
	}
	_, err = a.userRepo.GetUserByIdentifier(user.FirstName)
	if err == nil {
		return errors.New("user exists")
	}
	return a.userRepo.CreateUser(user)
}
