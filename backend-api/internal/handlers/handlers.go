package handlers

import "github.com/expose443/real-time-forum-golang/backend-api/pkg/logger"

type Client struct {
	logger *logger.LogLevel
}

func NewClient(logger *logger.LogLevel) *Client {
	return &Client{
		logger: logger,
	}
}
