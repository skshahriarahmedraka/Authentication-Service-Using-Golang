package handler

import "github.com/gin-gonic/gin"




func (H *DatabaseCollections)HandlerHealth() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(200, gin.H{"status": 200, "is_alive": true})
	}
}
