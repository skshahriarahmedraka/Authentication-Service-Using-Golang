package handler 

import (
	// "app/LogError"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/model"
	// "app/token"
	// "context"
	"fmt"
	// "log"
	"net/http"
	// "os"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/golang-jwt/jwt/v4"
	// "github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func (H *DatabaseCollections)Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var user model.UserData
		err := c.BindJSON(&user)
		if err != nil {
			// LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", user)
		
		fmt.Println("😍 Register Successfull")

		c.JSON(http.StatusOK,gin.H{"message":"successfully signed up"})
	}
}


