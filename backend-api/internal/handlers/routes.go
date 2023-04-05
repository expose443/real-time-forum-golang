package handlers

import (
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/middleware"
)

func (c *Client) SetupEndpoints() {
	router := http.NewServeMux()

	router.HandleFunc("/login", middleware.POST(c.Login))
	router.HandleFunc("/register", middleware.POST(c.Register))
}
