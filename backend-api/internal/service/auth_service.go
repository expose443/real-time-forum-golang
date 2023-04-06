package service

import "github.com/expose443/real-time-forum-golang/backend-api/internal/repository"

type AuthService struct {
	*repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo,
	}
}
