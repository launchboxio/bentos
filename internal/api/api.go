package api

import (
  "github.com/gin-gonic/gin"
  clusterv1 "github.com/launchboxio/bentos/internal/api/cluster/v1"
  functionv1 "github.com/launchboxio/bentos/internal/api/function/v1"
  instancev1 "github.com/launchboxio/bentos/internal/api/instance/v1"
  "net/http"
)

func New() (*gin.Engine, error) {
  r := gin.Default()

  r.GET("/health", health)

  apiRouter := r.Group("/api")
  instancev1.AddRoutes(apiRouter)
  functionv1.AddRoutes(apiRouter)
  clusterv1.AddRoutes(apiRouter)

  return r, nil
}

func health(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "status": "healthy",
  })
}
