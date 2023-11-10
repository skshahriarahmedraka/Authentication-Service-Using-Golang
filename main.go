package main

import (
	"context"
	// "errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/config"
	// "github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/init"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/routes"

	"github.com/gin-gonic/gin"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)
func init() {

	config.LoadEnvVars()
	config.LoadAdmin()

}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a Authentication-Service-Using-Golang.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
// func main() {
// 	//  SET MODE TO RELEASE
// 	gin.SetMode(gin.ReleaseMode)


// 	fmt.Println("🚀✨ Api is started ...")

// 	r := gin.New()
	

// 	r.Use(gin.Logger())

// 	routes.Routes(r)

// 	log.Println("Server is started in PORT ",os.Getenv("HOST_ADDR")," ...👨‍💻 ")
// 	if e := r.Run(os.Getenv("HOST_ADDR")); e != nil {
// 		log.Fatalln("❌ ERROR when Server is start   👨‍💻 : ", e)
// 	}

// }


func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())

	// Set up your routes
	routes.Routes(r)

	server := &http.Server{
		Addr:    os.Getenv("HOST_ADDR"),
		Handler: r,
	}

	go func() {
		fmt.Println("🚀✨ Api is started ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
	
	return nil
}
