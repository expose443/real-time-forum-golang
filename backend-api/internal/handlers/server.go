package handlers

import (
	"net/http"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/middleware"
)

func (app *Client) Server() *http.Server {
	router := http.NewServeMux()

	router.HandleFunc("/login", middleware.POST(middleware.Auth(app.Login)))
	router.HandleFunc("/register", middleware.POST((middleware.Auth(app.Register))))
	router.HandleFunc("/ws", app.WsHandler)

	return &http.Server{
		ReadTimeout:  time.Second * time.Duration(app.config.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(app.config.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(app.config.IdleTimeout),
		Addr:         app.config.Port,
		Handler:      router,
	}
}
