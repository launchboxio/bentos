package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	clusterGroup := router.Group("/clusters/v1")
	clusterGroup.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"route": "/api/clusters/v1",
		})
	})
}
