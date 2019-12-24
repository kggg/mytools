package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initroute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"msg":    "hello world",
		})

	})
	r.GET("/test", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"status": true, "msg":"this is a test page"})
	})
	return r
}
