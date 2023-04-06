package handlers

import (
	"github.com/expose443/real-time-forum-golang/backend-api/internal/service"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/config"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/logger"
)

type Client struct {
	AuthService service.AuthService
	Logger      *logger.LogLevel
	Config      *config.Config
}

func NewClient(s Client) *Client {
	return &Client{
		Logger:      s.Logger,
		AuthService: s.AuthService,
	}
}
