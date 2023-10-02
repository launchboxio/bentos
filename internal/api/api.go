package api

import (
  "github.com/gin-gonic/gin"
  instancev1 "github.com/launchboxio/bentos/internal/api/instance/v1"
  v1routes "github.com/launchboxio/bentos/internal/api/v1"
)

func New() (*gin.Engine, error) {
  r := gin.Default()

  apiRouter := r.Group("/api")
  instancev1.AddRoutes(apiRouter)

  v1Router := apiRouter.Group("/v1")
  v1routes.AddRoutes(v1Router)

  return r, nil
}
