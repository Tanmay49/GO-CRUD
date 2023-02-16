package Routes

import "github.com/gin-gonic/gin"

func Trial(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Worldsss",
	})
}
