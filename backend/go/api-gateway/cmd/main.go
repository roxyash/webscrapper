package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"webscrapper/api-gateway/internal/config"
	"webscrapper/api-gateway/internal/controller/httpv1"
	"webscrapper/api-gateway/internal/repository"
	"webscrapper/api-gateway/internal/service"
	"webscrapper/pkg/db/pg"
	"webscrapper/pkg/logging"
	httpserver "webscrapper/pkg/server/http"
)

var (
	logger = logging.GetLogger()
	conf   = config.GetConfig()
)

func main() {
	fmt.Println(conf)
	logger.Info("init application")

	// Init pg connection
	pgConn, err := pg.Connect(pg.Config{
		Host:     conf.Postgres.Host,
		Port:     conf.Postgres.Port,
		User:     conf.Postgres.User,
		Password: conf.Postgres.Password,
		Database: conf.Postgres.Database,
	})

	defer pg.Close(pgConn)

	if err != nil {
		logger.Fatalf("connect to postgres with err - %s", err.Error())
	}

	// Init dependencies
	repo := repository.New(pgConn)
	services := service.New(repo)

	// Init controllers
	httpRouter := httpv1.New(services)

	// Init servers
	httpServer := httpserver.New(httpRouter, conf.App.Port)

	// Start servers
	go func() {
		httpServer.Start()
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	httpServer.ListenQuitChanBlocking(quit)

	logger.Info("application started")
}
