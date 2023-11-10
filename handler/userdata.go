package handler

import (
	// "app/LogError"
	// "app/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/models"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/pkg/authentication"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) UserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Login successful !!! ")
		//c.Param("profileid")

		authToken, err := c.Cookie("Auth1")
		refreshToken, err := c.Cookie("Refresh")
		var claims *model.TokenClaims

		if authToken == "" {
			if refreshToken == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User"})
				c.Abort()
				return
			} else {
				refreshClaims, str := authentication.ValidateRefreshJWT(refreshToken)
				if str != "" {
					fmt.Println("🚀 unauthenticated user: ", errors.New("jwt error"))
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
						fmt.Println("❌🔥 error in StatusInternalServerError token generation  ", err)
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
						fmt.Println("❌🔥 error in StatusInternalServerError token generation  ", err)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

						return
					}

					c.SetCookie("Auth1", tokenString, 60*60*24, "/", os.Getenv("DOMAIN_ADDR"), false, true)
					c.SetCookie("refresh", tokenString2, 60*60*24*2, "/", os.Getenv("DOMAIN_ADDR"), false, true)
				}

			}

			// c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User"})
			// c.Abort()
			// return
		}else {
			str := ""
			claims, str = authentication.ValidateJWT(authToken)
			if str != "" {
				fmt.Println("🚀 ~ Auth1 : ", errors.New("jwt error"))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth1 error"})
				c.Abort()
				return
				
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var UserDataDB model.UserData
		//options.FindOne()
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err = H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"_id": objectID}).Decode(&UserDataDB)

		if claims.AccountType == "admin" || claims.Email == UserDataDB.Email{

			if err != nil {
				fmt.Println("profiledata mongo find one error", err)
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			} else {
				// objID, _ := primitive.ObjectIDFromHex("")
				// UserDataDB.ID = objID
				UserDataDB.Password = ""
				c.JSON(http.StatusOK, UserDataDB)
			}
		}else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		}

	}
}
