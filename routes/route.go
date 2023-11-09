
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
	r.POST("/user/login", H.Login())
	r.POST("/user/register", H.Register())
	
}
