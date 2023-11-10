package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/pkg/monodb"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func Routes(r *gin.Engine) {
	mongoConn := monodb.MongodbConnection()
	H := monodb.DatabaseInitialization(mongoConn)

	r.GET("/health", H.HandlerHealth())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/public", H.Home())
	r.POST("/login", H.Login())
	r.POST("/register", H.Register())
	r.GET("/logout", H.Logout())
	r.GET("/:id", H.UserData())
	r.GET("/alluser", H.AllUserData())
}
