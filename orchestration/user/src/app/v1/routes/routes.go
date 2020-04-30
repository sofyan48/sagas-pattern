package routes

import (
	"github.com/gin-gonic/gin"
	health "github.com/sofyan48/user/src/app/v1/api/health/controller"
	order "github.com/sofyan48/user/src/app/v1/api/order/controller"
	payment "github.com/sofyan48/user/src/app/v1/api/payment/controller"
	user "github.com/sofyan48/user/src/app/v1/api/user/controller"
	"github.com/sofyan48/user/src/middleware"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	User       user.UserControllerInterface
	Order      order.OrderControllerInterface
	Payment    payment.PaymentControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health:  health.HealthControllerHandler(),
		User:    user.UserControllerHandler(),
		Order:   order.OrderControllerHandler(),
		Payment: payment.PaymentControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initUser(router)
	rLoader.initOrder(router)
	rLoader.initPayment(router)
}
