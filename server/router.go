package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-starter/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(handlers.HealthController)
	router.GET("/health", health.Status)

	return router

}
