package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Public Home  . . . ")
	}
}
