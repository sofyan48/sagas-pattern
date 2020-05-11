package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_auth/src/app/v1/api/login/service"
	"github.com/sofyan48/svc_auth/src/utils/rest"
)

// LoginController ...
type LoginController struct {
	Service service.LoginServiceInterface
}

// LoginControllerHandler ...
func LoginControllerHandler() *LoginController {
	return &LoginController{
		Service: service.LoginServiceHandler(),
	}
}

// LoginControllerInterface ...
type LoginControllerInterface interface {
	PostLoginClient(context *gin.Context)
}

// PostLoginClient ...
func (handler *LoginController) PostLoginClient(context *gin.Context) {
	payload := &entity.LoginRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := handler.Service.PostLogin(payload, context)

	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
}
