package apis

import "github.com/gin-gonic/gin"

// Hello hello
func Hello(c *gin.Context) {
	c.String(200, "data")
}
