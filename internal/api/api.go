package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() (*gin.Engine, error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r, nil
}
