package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/handlers"
	"github.com/expose443/real-time-forum-golang/backend-api/internal/repository"
	"github.com/expose443/real-time-forum-golang/backend-api/internal/service"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/config"
	"github.com/expose443/real-time-forum-golang/backend-api/pkg/logger"
)

func main() {
	var configFile string = "config.json"

	logger := logger.NewLog()
	cfg, err := config.Init(configFile)
	if err != nil {
		logger.Error.Fatal(err)
		return
	}
	logger.Warning.Printf("using %s file for set up", configFile)

	db, err := repository.NewSqliteDB(&cfg.DB)
	if err != nil {
		logger.Error.Fatal(err)
	}
	defer db.Close()
	logger.Info.Printf("database is ready")

	dao := repository.NewDao(db)
	authService := service.NewAuthService(dao)

	app := handlers.NewClient(handlers.Services{
		Logger:      logger,
		AuthService: authService,
		Config:      cfg,
	})

	server := app.Server()

	go func() {
		logger.Info.Printf("server started at %s", cfg.Address)
		if err := server.ListenAndServe(); err != nil {
			logger.Error.Print(err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info.Print("shutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		logger.Error.Printf("server shutdown: %s", err)
	} else {
		logger.Info.Print("server gracefully stoped")
	}
}
