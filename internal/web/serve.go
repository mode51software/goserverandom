package web

import (
	"github.com/gin-gonic/gin"
)

const ()

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	//r.Use(ConfigureResponse())

	//r.GET("/ping", setupPingPong)
	r.GET("/genrandomfile/:datasz", genRandomFile)

	return r
}

func ConfigureResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/octet-string")
		c.Next()
	}
}
