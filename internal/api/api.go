package api

import (
	"github.com/gin-gonic/gin"
	v1routes "github.com/launchboxio/bentos/internal/api/v1"
)

func New() (*gin.Engine, error) {
	r := gin.Default()

	router := r.Group("/api/v1")
	v1routes.AddRoutes(router)

	return r, nil
}
