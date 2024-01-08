package service

import (
	"webscrapper/api-gateway/internal/model"
	"webscrapper/api-gateway/internal/repository"
)

type Auth interface {
	Login(req model.LoginRequest) (model.LoginResponse, error)
	Register(req model.RegisterRequest) (model.RegisterResponse, error)
}

type authService struct {
	repo *repository.Repository
}

func NewAuth(repo *repository.Repository) Auth {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(req model.LoginRequest) (model.LoginResponse, error) {
	return s.repo.WebApi.Internal.Auth.Login(req)
}

func (s *authService) Register(req model.RegisterRequest) (model.RegisterResponse, error) {
	return s.repo.WebApi.Internal.Auth.Register(req)
}
