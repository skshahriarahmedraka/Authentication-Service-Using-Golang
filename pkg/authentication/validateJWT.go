package authentication

import (
	"os"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/models"
)

func ValidateJWT(s string) (claims *model.TokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(s, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}

	return claims, msg

}

func ValidateRefreshJWT(s string) (claims *model.TokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(s, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("COOKIE_SECRET_JWT_AUTH2")), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}

	return claims, msg

}
