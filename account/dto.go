package account

import "github.com/golang-jwt/jwt/v5"

type AccountDto struct {
	Id       int    `json:"id"`
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
	Token string `json:"token"`
}
