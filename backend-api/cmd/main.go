package main

import (
	"net/http"

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

	repository := repository.NewRepository(db)

	authService := service.NewAuthService(repository.UserRepository)

	app := handlers.NewClient(logger)

	router := http.NewServeMux()
	app.SetupEndpoints(router)
	logger.Error.Fatal(http.ListenAndServe(":8080", router))
}
