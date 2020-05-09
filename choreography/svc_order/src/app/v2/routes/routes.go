package routes

import (
	"github.com/gin-gonic/gin"
	order "github.com/sofyan48/svc_order/src/app/v2/api/order/controller"
	"github.com/sofyan48/svc_order/src/middleware"
)

// VERSION ...
const VERSION = "v2"

// V2RouterLoader types
type V2RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Order      order.OrderControllerInterface
}

// V2RouterLoaderHandler ...
func V2RouterLoaderHandler() *V2RouterLoader {
	return &V2RouterLoader{
		Order: order.OrderControllerHandler(),
	}
}

// V2RouterLoaderInterface ...
type V2RouterLoaderInterface interface {
	V2Routes(router *gin.Engine)
}

// V2Routes Params
// @router: gin.Engine
func (rLoader *V2RouterLoader) V2Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initOrder(router)
}
