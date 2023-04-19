package handlers

import (
	"net/http"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/middle"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *Client) Server() *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.HandleFunc("/login", middle.POST(middle.Auth(app.Login)))
	router.HandleFunc("/register", middle.POST((middle.Auth(app.Register))))
	router.HandleFunc("/ws", app.WsHandler)

	return &http.Server{
		ReadTimeout:  time.Second * time.Duration(app.config.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(app.config.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(app.config.IdleTimeout),
		Addr:         app.config.Port,
		Handler:      router,
	}
}
