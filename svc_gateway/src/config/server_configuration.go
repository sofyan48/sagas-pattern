package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/nemo/src/middleware"
)

// SetupRouter server router configuration
func SetupRouter() *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}
