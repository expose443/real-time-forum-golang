package handlers

import (
	"github.com/expose443/real-time-forum-golang/backend-api/internal/service"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/config"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/logger"
)

type Client struct {
	authService service.AuthService
	logger      *logger.LogLevel
	config      *config.Config
}

type Services struct {
	AuthService service.AuthService
	Logger      *logger.LogLevel
	Config      *config.Config
}

func NewClient(s Services) *Client {
	return &Client{
		logger:      s.Logger,
		authService: s.AuthService,
		config:      s.Config,
	}
}
