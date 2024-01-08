package internal

import (
	"context"
	"fmt"
	"github.com/roxyash/webscrapper_proto/gen/go/auth"
	"google.golang.org/grpc"
	"webscrapper/api-gateway/internal/model"
)

type Auth interface {
	Login(req model.LoginRequest) (model.LoginResponse, error)
	Register(req model.RegisterRequest) (model.RegisterResponse, error)
	RefreshToken(req model.RefreshTokenRequest) (model.RefreshTokenResponse, error)
	VerifyAccessToken(req model.VerifyTokenRequest) (model.VerifyTokenResponse, error)
}

type authRepo struct {
	authClient auth.AuthServiceClient
}

func NewAuth() Auth {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.Auth.Host, conf.Auth.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("connection to auth service failed: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Fatalf("close connection for auth service: %v", err)
		}
	}(conn)

	authClient := auth.NewAuthServiceClient(conn)

	return &authRepo{
		authClient: authClient,
	}
}

func (r *authRepo) Login(req model.LoginRequest) (model.LoginResponse, error) {
	resp, err := r.authClient.Login(context.Background(), &auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{
		JWT: model.JWT{
			AccessToken:  resp.Jwt.AccessToken,
			RefreshToken: resp.Jwt.RefreshToken,
		},
	}, nil
}

func (r *authRepo) Register(req model.RegisterRequest) (model.RegisterResponse, error) {
	resp, err := r.authClient.UserRegister(context.Background(), &auth.UserRegisterRequest{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsAdmin:   req.IsAdmin,
		IsArchive: req.IsArchive,
	})

	if err != nil {
		return model.RegisterResponse{}, err
	}

	return model.RegisterResponse{
		JWT: model.JWT{
			AccessToken:  resp.Jwt.AccessToken,
			RefreshToken: resp.Jwt.RefreshToken,
		},
	}, nil
}

func (r *authRepo) RefreshToken(req model.RefreshTokenRequest) (model.RefreshTokenResponse, error) {
	resp, err := r.authClient.RefreshToken(context.Background(), &auth.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		return model.RefreshTokenResponse{}, err
	}

	return model.RefreshTokenResponse{
		JWT: model.JWT{
			AccessToken:  resp.Jwt.AccessToken,
			RefreshToken: resp.Jwt.RefreshToken,
		},
	}, nil
}

func (r *authRepo) VerifyAccessToken(req model.VerifyTokenRequest) (model.VerifyTokenResponse, error) {
	resp, err := r.authClient.VerifyAccessToken(context.Background(), &auth.VerifyTokenRequest{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return model.VerifyTokenResponse{}, err
	}

	return model.VerifyTokenResponse{
		UserID:  resp.UserId,
		IsAdmin: resp.IsAdmin,
	}, nil
}
