package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	instanceGroup := router.Group("/instances/v1")
	instanceGroup.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"route": "/api/instances/v1",
		})
	})
}
