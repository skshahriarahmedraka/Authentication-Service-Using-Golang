package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Email       string `json:"email"`
	AccountType string `json:"accounttype"`
	jwt.StandardClaims
}
