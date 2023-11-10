package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func (H *DatabaseCollections) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var loginUser model.LoginModel
		var dbUser model.UserData

		err := c.BindJSON(&loginUser)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		fmt.Println("🚀", loginUser)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		err = H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"email": loginUser.Email}).Decode(&dbUser)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
			return
		}

		userPass := []byte(loginUser.Password)
		dbPass := []byte(dbUser.Password)
		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
		if passErr != nil {
			logger.Error().Msg("❌🔥 Error message :" + passErr.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//  AUTH1 GENEGRATION
		expirationTime := time.Now().Add(time.Hour * 1000)
		claims := &model.TokenClaims{
			Email:       dbUser.Email,
			AccountType: dbUser.AccountType,

			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")))

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

			return
		}

		//  noAUTH1 GENEGRATION
		expirationTime2 := time.Now().Add(time.Hour * 1000)
		claims2 := &model.TokenClaims{
			Email:       dbUser.Email,
			AccountType: dbUser.AccountType,

			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime2.Unix(),
			},
		}

		token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)

		tokenString2, err := token2.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH2")))

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})
			return
		}

		c.SetCookie("Auth1", tokenString, 60*60*24, "/", os.Getenv("DOMAIN_ADDR"), false, true)
		c.SetCookie("Refresh", tokenString2, 60*60*24*2, "/", os.Getenv("DOMAIN_ADDR"), false, true)

		fmt.Println("😍 Login Successfull")
		c.JSON(http.StatusOK, gin.H{"id": dbUser.ID.Hex(), "message": "Login Successfull"})

	}
}
