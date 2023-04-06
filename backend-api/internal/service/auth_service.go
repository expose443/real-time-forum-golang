package service

import "github.com/expose443/real-time-forum-golang/backend-api/internal/repository"

type AuthService interface{}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		userRepo: dao.NewUserRepo(),
	}
}

type authService struct {
	userRepo repository.UserRepository
}
