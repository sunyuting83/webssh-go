package router

import (
	"webssh/controller/apis"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/index", apis.Hello)
		api.GET("/websocket", apis.WsPage)
	}

	return router
}
