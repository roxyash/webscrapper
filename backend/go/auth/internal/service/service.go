package service

import (
	"webscrapper/auth/internal/config"
	"webscrapper/auth/internal/repository"
)

type Service struct {
	Auth
}

var (
	conf = config.GetConfig()
)

func New(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuth(repo),
	}
}
