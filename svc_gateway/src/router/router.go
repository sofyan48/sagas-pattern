package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sofyan48/nemo/src/app/v1/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	version1 := v1.V1RouterLoaderHandler()
	version1.V1Routes(routers)
}
