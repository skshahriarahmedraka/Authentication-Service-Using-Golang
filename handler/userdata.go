package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/models"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/pkg/authentication"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"time"
)

func (H *DatabaseCollections) UserData() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken, _ := c.Cookie("Auth1")
		refreshToken, _ := c.Cookie("Refresh")
		var claims *model.TokenClaims
		if authToken == "" {
			if refreshToken == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User"})
				c.Abort()
				return
			} else {
				refreshClaims, str := authentication.ValidateRefreshJWT(refreshToken)
				if str != "" {
					logger.Error().Msg("❌🔥 Error message :" + errors.New("unauthenticated user").Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User"})
					c.Abort()
					return

				} else {
					claims = refreshClaims
					expirationTime := time.Now().Add(time.Hour * 100)
					newclaims := &model.TokenClaims{
						Email:       refreshClaims.Email,
						AccountType: refreshClaims.AccountType,
						StandardClaims: jwt.StandardClaims{
							ExpiresAt: expirationTime.Unix(),
						},
					}

					token := jwt.NewWithClaims(jwt.SigningMethodHS256, newclaims)

					tokenString, err := token.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")))

					if err != nil {
						logger.Error().Msg("❌🔥 Error message :" + err.Error())
						c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})
						return
					}

					//  noAUTH1 GENEGRATION
					expirationTime2 := time.Now().Add(time.Hour * 1000)
					newclaims2 := &model.TokenClaims{
						Email:       refreshClaims.Email,
						AccountType: refreshClaims.AccountType,
						StandardClaims: jwt.StandardClaims{
							ExpiresAt: expirationTime2.Unix(),
						},
					}

					token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, newclaims2)

					tokenString2, err := token2.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH2")))

					if err != nil {
						logger.Error().Msg("❌🔥 Error message :" + err.Error())

						c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

						return
					}

					c.SetCookie("Auth1", tokenString, 60*60*24, "/", os.Getenv("DOMAIN_ADDR"), false, true)
					c.SetCookie("Refresh", tokenString2, 60*60*24*2, "/", os.Getenv("DOMAIN_ADDR"), false, true)
				}

			}
		} else {
			str := ""
			claims, str = authentication.ValidateJWT(authToken)
			if str != "" {
				logger.Error().Msg("❌🔥 Error message :" + errors.New("unauthenticated user using JWT").Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthenticated user using JWT"})
				c.Abort()
				return
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var UserDataDB model.UserData
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"_id": objectID}).Decode(&UserDataDB)
		if claims.AccountType == "admin" || claims.Email == UserDataDB.Email {
			if err != nil {
				logger.Error().Msg("❌🔥 Error message :" + err.Error())
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			} else {

				UserDataDB.Password = ""
				c.JSON(http.StatusOK, UserDataDB)
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		}

	}
}
