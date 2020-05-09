package router

import (
	"github.com/gin-gonic/gin"
	v2 "github.com/sofyan48/svc_payment/src/app/v2/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	version2 := v2.V2RouterLoaderHandler()
	version2.V2Routes(routers)
}
