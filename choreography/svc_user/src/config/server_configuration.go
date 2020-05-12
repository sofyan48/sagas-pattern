package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sofyan48/svc_user/src/middleware"
)

// SetupEngine server router configuration
func SetupEngine(env string) *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddlewareHandler()
	configEnvironment(env)
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}

// ConfigEnvironment ...
func configEnvironment(env string) {
	switch env {
	case "development":
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	case "production":
		gin.SetMode(gin.ReleaseMode)
		log.Println("Engine Running")
	default:
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

}
