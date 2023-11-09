package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/init"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/router"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/routes"

	"github.com/gin-gonic/gin"
)

func init() {

	init.LoadEnvVars()

}

func main() {
	//  SET MODE TO RELEASE
	gin.SetMode(gin.ReleaseMode)


	fmt.Println("🚀✨ Api is started ...")

	r := gin.New()
	
	// r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	


	routes.Routes(r)

	log.Println("Server is started in PORT ",os.Getenv("HOST_ADDR")," ...👨‍💻 ")
	if e := r.Run(os.Getenv("HOST_ADDR")); e != nil {
		log.Fatalln("❌ ERROR when Server is start   👨‍💻 : ", e)
	}

}
