package di

import (
	"tasks/internal/adapters/repository"
	"tasks/internal/adapters/transport"
	"tasks/internal/config"
	"tasks/internal/db"
	"tasks/internal/service"
)

func Initialize() (*transport.Router, string, error) {
	cfg := config.LoadConfig()

	dbConnString := cfg.GetDBConnectionString()

	dbConn, err := db.Connect(dbConnString)
	if err != nil {
		return nil, "", err
	}

	taskRepo := repository.NewPostgresRepo(dbConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := transport.NewTaskHandler(taskService)

	router := transport.NewRouter(taskHandler)

	return router, cfg.Port, nil
}
