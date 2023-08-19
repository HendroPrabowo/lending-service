package account

import "github.com/golang-jwt/jwt/v5"

type AccountDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Claims struct {
	Account
	jwt.RegisteredClaims
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}
