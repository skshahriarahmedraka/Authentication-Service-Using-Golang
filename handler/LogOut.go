package handler

import (
	// "html/template"
	

	"github.com/gin-gonic/gin"
	// "time"
	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
	"github.com/gin-contrib/sessions"
)

func (H *DatabaseCollections)Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Clear the user session
		session.Clear()
		session.Save()
	
		c.JSON(200, gin.H{"message": "Logout successful"})

	// http.Redirect(w,r,"/login",http.StatusSeeOther)
	}
}