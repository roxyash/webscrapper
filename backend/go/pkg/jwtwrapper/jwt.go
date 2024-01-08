package jwtwrapper

import (
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	key []byte
}

func New(jwtSecretKey string) *Jwt {
	return &Jwt{
		key: []byte(jwtSecretKey),
	}
}

func (j *Jwt) GenerateTokens(claims jwt.MapClaims, refreshClaims jwt.MapClaims) (string, string, error) {
	accessTokenString, err := j.GenerateAccessToken(claims)
	if err != nil {
		return "", "", nil
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(j.key)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (j *Jwt) GenerateAccessToken(claims jwt.MapClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(j.key)

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (j *Jwt) ValidateToken(tokenString string) (jwt.Claims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
