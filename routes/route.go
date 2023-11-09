
package routes

import (
	
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/pkg/monodb"

)

func Routes(r *gin.Engine) {
	// fmt.Println("🚀 ~ file: routers.go ~ line 21 ~ funcUserRoutes ~ database.Client : ", database.Client)
	DB2 := monodb.MongodbConnection()
	H := monodb.DatabaseInitialization(DB2)
	fmt.Println("🚀 ~ file: WithoutAuth.go ~ line 13 ~ funcRouteWithoutAuth ~ RouteWithoutAuth  ", H)
	
	r.GET("/", H.Home())
	r.POST("/login", H.Login())
	r.POST("/register", H.Register())
	r.GET("/logout", H.Logout())
	// r.GET("/:id", H.UserProfileData())
	// r.POST("/update", H.Update())
	// r.POST("/delete", H.Delete())
	// r.GET("/alluser", H.AllUser())
	// r.GET("/admindashboard", H.AdminDashboard())
}
