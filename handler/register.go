package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/config"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var Validate = validator.New()

func (H *DatabaseCollections) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
		var user model.UserData
		var regUser model.RegModel
		err := c.BindJSON(&regUser)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", regUser)
		if _, b := config.AdminEmails[regUser.Email]; b {
			user.AccountType = "admin"

		} else {
			user.AccountType = "normal"
		}

		user.ID = primitive.NewObjectID()
		user.Firstname = regUser.Firstname
		user.Lastname = regUser.Lastname
		user.Email = regUser.Email
		user.Address = ""
		user.Telephone = ""
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		count, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).CountDocuments(ctx, bson.M{"Email": user.Email})
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "already registered user")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(regUser.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
		}

		res, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).InsertOne(ctx, user)
		if err == nil {
			logger.Info().Msg("📢 Info message :" + "successfully registered user")
		}
		_ = res

		//  AUTH1 token GENEGRATION
		expirationTime := time.Now().Add(time.Hour * 100)
		claims := &model.TokenClaims{
			Email:       user.Email,
			AccountType: user.AccountType,

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

		//  Refresh token GENEGRATION
		expirationTime2 := time.Now().Add(time.Hour * 1000)
		claims2 := &model.TokenClaims{
			Email:       user.Email,
			AccountType: user.AccountType,

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
		fmt.Println("😍 Register Successfull")

		c.JSON(http.StatusOK, gin.H{"message": "successfully signed up"})
	}
}
