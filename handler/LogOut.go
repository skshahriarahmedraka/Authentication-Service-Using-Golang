package handler

import (
	// "html/template"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "time"
	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
	// "github.com/gin-contrib/sessions"
)

func (H *DatabaseCollections)Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		c.SetCookie("Auth1","",-1,"/",os.Getenv("DOMAIN_ADDR"),false , true)
		c.SetCookie("Refresh","",-1,"/",os.Getenv("DOMAIN_ADDR"),false , true)

		c.Redirect(http.StatusTemporaryRedirect, "/public")
	
	}
}