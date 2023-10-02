package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutes(router *gin.RouterGroup) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"route": "/api/instance/v1",
		})
	})
}
