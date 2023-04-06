package handlers

import (
	"net/http"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/middleware"
)

func (app *Client) Server() *http.Server {
	router := http.NewServeMux()

	router.HandleFunc("/login", middleware.POST(app.Login))
	router.HandleFunc("/register", middleware.POST(app.Register))

	return &http.Server{
		ReadTimeout:  time.Second * time.Duration(app.Config.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(app.Config.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(app.Config.IdleTimeout),
		Addr:         app.Config.Port,
		Handler:      router,
	}
}
