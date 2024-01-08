package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/roxyash/webscrapper_proto/gen/go/auth"
	"golang.org/x/crypto/bcrypt"
	"time"
	"webscrapper/auth/internal/model"
	"webscrapper/auth/internal/repository"
	"webscrapper/auth/internal/repository/cache"
	"webscrapper/pkg/jwtwrapper"
)

type Auth interface {
	Login(req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(req *auth.UserRegisterRequest) (*auth.UserRegisterResponse, error)
	RefreshToken(req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	VerifyAccessToken(req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error)
}

type authService struct {
	repo      *repository.Repository
	jwtVerify *jwtwrapper.Jwt
}

func NewAuth(repo *repository.Repository) Auth {
	jwtVerify := jwtwrapper.New(conf.Other.JwtSecretKey)

	return &authService{
		repo:      repo,
		jwtVerify: jwtVerify,
	}
}

func (s *authService) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// Check user in cache
	userCache, err := s.repo.Cache.Redis.User.GetByUsername(req.Username)
	if err != nil {
		switch err.Error() {
		case cache.ErrCacheMiss.Error():
			// Check in db
			user, err := s.repo.Storage.Postgres.User.GetByEmailOrUsername(req.Username)
			if err != nil {
				return nil, err
			}

			// Set in cache
			err = s.repo.Cache.Redis.User.Set(user)
			if err != nil {
				return nil, err
			}

			userCache = user
		default:
			return nil, err
		}
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(userCache.Password), []byte(req.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return nil, errors.New("password does not match")
	}

	// Generate access token claims
	accessClaims := jwt.MapClaims{
		"user_id":  userCache.ID.String(),
		"is_admin": userCache.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	// Generate refresh token claims
	refreshClaims := jwt.MapClaims{
		"user_id":  userCache.ID.String(),
		"is_admin": userCache.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Generate refresh and access tokens
	accessToken, refreshToken, err := s.jwtVerify.GenerateTokens(accessClaims, refreshClaims)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{Jwt: &auth.JWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}}, nil

}

func (s *authService) Register(req *auth.UserRegisterRequest) (*auth.UserRegisterResponse, error) {
	// Check if user already exists
	_, err := s.repo.Storage.Postgres.User.GetByEmailOrUsername(req.Username)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user, err := s.repo.Storage.Postgres.User.Create(model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsArchive: req.IsArchive,
		IsAdmin:   req.IsAdmin,
	})
	if err != nil {
		return nil, err
	}

	// Set in cache
	err = s.repo.Cache.Redis.User.Set(user)
	if err != nil {
		return nil, err
	}

	// Generate access token claims
	accessClaims := jwt.MapClaims{
		"user_id":  user.ID.String(),
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	// Generate refresh token claims
	refreshClaims := jwt.MapClaims{
		"user_id":  user.ID.String(),
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Generate refresh and access tokens
	accessToken, refreshToken, err := s.jwtVerify.GenerateTokens(accessClaims, refreshClaims)
	if err != nil {
		return nil, err
	}

	return &auth.UserRegisterResponse{
		Jwt: &auth.JWT{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *authService) RefreshToken(req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	// Check valid RefreshToken
	claims, err := s.jwtVerify.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Conversion type to jwt.MapClaims
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claims to MapClaims")
	}

	// Generate new AccessToken
	newClaims := jwt.MapClaims{
		"user_id":  mapClaims["user_id"],
		"is_admin": mapClaims["is_admin"],
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	newAccessToken, err := s.jwtVerify.GenerateAccessToken(newClaims)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenResponse{Jwt: &auth.JWT{
		AccessToken:  newAccessToken,
		RefreshToken: req.RefreshToken,
	}}, nil
}

func (s *authService) VerifyAccessToken(req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	// Validate the token
	claims, err := s.jwtVerify.ValidateToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	// Convert claims to jwt.MapClaims
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claims to MapClaims")
	}

	// Get the user ID from the claims
	userID, ok := mapClaims["user_id"].(string)
	if !ok {
		return nil, errors.New("cannot get user_id from token claims")
	}

	// Get the user from the database
	user, err := s.repo.Storage.Postgres.User.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &auth.VerifyTokenResponse{
		UserId:  user.ID.String(),
		IsAdmin: user.IsAdmin,
	}, nil
}
