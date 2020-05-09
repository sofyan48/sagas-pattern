package routes

import (
	"github.com/gin-gonic/gin"

	payment "github.com/sofyan48/svc_payment/src/app/v2/api/payment/controller"
	"github.com/sofyan48/svc_payment/src/middleware"
)

// VERSION ...
const VERSION = "v2"

// V2RouterLoader types
type V2RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Payment    payment.PaymentControllerInterface
}

// V2RouterLoaderHandler ...
func V2RouterLoaderHandler() *V2RouterLoader {
	return &V2RouterLoader{
		Payment: payment.PaymentControllerHandler(),
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
	rLoader.initPayment(router)
}
