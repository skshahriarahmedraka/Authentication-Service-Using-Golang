package handler

import (
	// "app/LogError"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/model"
	// "app/token"
	// "context"
	"fmt"
	"net/http"
	// "os"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v4"
	// "go.mongodb.org/mongo-driver/bson"
	// "golang.org/x/crypto/bcrypt"
)

func (H *DatabaseCollections) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var user model.LoginModel
		// var dbUser model.UserData

		err := c.BindJSON(&user)
		if err != nil {
			// LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		fmt.Println("🚀", user)

		fmt.Println("😍 Login Successfull")
		c.JSON(http.StatusOK, user)

	}
}
