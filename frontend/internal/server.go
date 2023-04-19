package internal

import (
	"net/http"
	"time"

	"github.com/expose443/real-time-forum-golang/frontend/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Server() *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./templates/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.HandleFunc("/", handlers.WebSocket)
	router.HandleFunc("/sign-in", handlers.Signin)
	router.HandleFunc("/sign-up", handlers.Signup)
	return &http.Server{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Addr:         ":8080",
		Handler:      router,
	}
}
