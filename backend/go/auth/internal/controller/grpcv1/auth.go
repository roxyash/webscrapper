package grpcv1

import (
	"context"
	"github.com/roxyash/webscrapper_proto/gen/go/auth"
	"webscrapper/auth/internal/service"
)

type authController struct {
	auth.UnimplementedAuthServiceServer
	s service.Auth
}

func NewAuth(s service.Auth) auth.AuthServiceServer {
	return &authController{s: s}
}

func (h *authController) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// Validate request

	// Call service
	resp, err := h.s.Login(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *authController) Register(ctx context.Context, req *auth.UserRegisterRequest) (*auth.UserRegisterResponse, error) {
	// Validate request

	// Call service
	resp, err := h.s.Register(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *authController) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	// Validate request

	// Call service
	resp, err := h.s.RefreshToken(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *authController) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	// Validate request

	// Call service
	resp, err := h.s.VerifyAccessToken(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
