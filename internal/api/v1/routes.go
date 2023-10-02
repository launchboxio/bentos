package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	instanceGroup := router.Group("/instance")
	instanceGroup.GET("/health", health)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
