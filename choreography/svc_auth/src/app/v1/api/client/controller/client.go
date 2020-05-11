package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/entity"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/service"
	"github.com/sofyan48/svc_auth/src/utils/rest"
)

// ClientController ...
type ClientController struct {
	Service service.ClientServiceInterface
}

// ClientControllerHandler ...
func ClientControllerHandler() *ClientController {
	return &ClientController{
		Service: service.ClientServiceHandler(),
	}
}

// ClientControllerInterface ...
type ClientControllerInterface interface {
	PostClient(context *gin.Context)
}

// PostClient ..
func (ctrl *ClientController) PostClient(context *gin.Context) {
	payload := &entity.ClientRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := ctrl.Service.CreateClient(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}
