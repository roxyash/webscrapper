package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWT JWT `json:"jwt"`
}

type RegisterRequest struct {
	Username  string `json:"Username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsArchive bool   `json:"is_active"`
	IsAdmin   bool   `json:"is_admin"`
}

type RegisterResponse struct {
	JWT JWT `json:"jwt"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	JWT JWT `json:"jwt"`
}

type VerifyTokenRequest struct {
	AccessToken string `json:"access_token"`
}

type VerifyTokenResponse struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
}
