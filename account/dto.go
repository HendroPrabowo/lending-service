package account

import "github.com/golang-jwt/jwt/v5"

type AccountDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type AccountListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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
	Name  string `json:"name"`
	Id    int    `json:"id"`
}
