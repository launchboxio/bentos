package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	functionGroup := router.Group("/functions/v1")
	functionGroup.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"route": "/api/functions/v1",
		})
	})
}
