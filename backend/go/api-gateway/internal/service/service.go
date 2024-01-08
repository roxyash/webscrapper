package service

import (
	"webscrapper/api-gateway/internal/config"
	"webscrapper/api-gateway/internal/repository"
	"webscrapper/pkg/logging"
)

type Service struct {
	Auth
}

var (
	conf   = config.GetConfig()
	logger = logging.GetLogger()
)

func New(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuth(repo),
	}
}
