package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	router.GET("/health", health)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
