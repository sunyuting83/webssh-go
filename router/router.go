package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/index", func(c *gin.Context) {
			c.String(200, "data")
		})
	}

	return router
}
