package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/user/src/app/v1/api/health/service"
	"github.com/sofyan48/user/src/app/v1/utility/rest"
)

// HealthController types
type HealthController struct {
	Service service.HealthServiceInterface
}

// HealthControllerHandler ...
func HealthControllerHandler() *HealthController {
	return &HealthController{
		Service: service.HealthServiceHandler(),
	}
}

// HealthControllerInterface ...
type HealthControllerInterface interface {
	Health(context *gin.Context)
}

// Health params
// @contex: gin Context
func (ctrl *HealthController) Health(context *gin.Context) {
	data := ctrl.Service.HealthService()
	rest.ResponseData(context, http.StatusOK, data)
	return
}
