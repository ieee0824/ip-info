package root

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"ip":         c.ClientIP(),
		"user-agent": c.Request.UserAgent(),
	})
}
