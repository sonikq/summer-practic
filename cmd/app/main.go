package main

import (
	"context"
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/handlers"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/servers/http"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/services"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Configurations
	config := cfg.Load("configs/app/.env")

	// Logger
	log := logger.New("debug", "skdf-abac-go-logger")
	defer func() {
		err := logger.CleanUp(log)
		log.Fatal("failed to cleanup logs", logger.Error(err))
	}()

	_db, err := db.ConnectDatabase(config.DB)
	if err != nil {
		log.Fatal("failed to initialize db", logger.Error(err))
	}

	repo := repositories.NewRepository(_db)

	service := services.NewService(repo)

	router := handlers.NewRouter(handlers.Option{
		Conf:    config,
		Logger:  log,
		Service: service,
	})

	server := http.NewServer(config.HTTP.Port, router)

	go func() {
		if err = server.Run(); err != nil {
			log.Fatal("failed to run http server", logger.Error(err))
			panic(err)
		}
	}()

	log.Info("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to stop server", logger.Error(err))
	}

}
