package main

import (
	"os"
	"os/signal"
	"syscall"
	"webscrapper/auth/internal/config"
	"webscrapper/auth/internal/controller/grpcv1"
	"webscrapper/auth/internal/repository"
	"webscrapper/auth/internal/service"
	"webscrapper/pkg/db/pg"
	"webscrapper/pkg/db/rd"
	"webscrapper/pkg/logging"
	grpcserver "webscrapper/pkg/server/grpc"
)

var (
	conf   = config.GetConfig()
	logger = logging.GetLogger()
)

func main() {
	logger.Infof("init application")

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

	// Init redis connection
	rdConn, err := rd.Connect(rd.Config{
		Host:     conf.Redis.Host,
		Port:     conf.Redis.Port,
		Username: conf.Redis.User,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	defer rd.Close(rdConn)

	if err != nil {
		logger.Fatalf("connect to redis with err - %s", err.Error())
	}

	// Init dependencies
	repo := repository.New(pgConn, rdConn)
	services := service.New(repo)

	// Init controllers
	grpcRouter := grpcv1.New(services)

	// Init servers
	grpcServer, err := grpcserver.New(conf.App.Port, grpcRouter)
	if err != nil {
		logger.Fatalf("failed to create gRPC server: %s", err)
	}

	// Start servers
	go func() {
		grpcServer.Start()
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	grpcServer.ListenQuitChanBlocking(quit)

	logger.Info("application started")
}
