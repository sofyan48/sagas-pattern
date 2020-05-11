package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/authorization/service"
	"github.com/sofyan48/svc_auth/src/utils/rest"
)

// AuthorizationController ...
type AuthorizationController struct {
	Service service.AuthorizationServiceInterface
}

// AuthorizationControllerHandler ...
func AuthorizationControllerHandler() *AuthorizationController {
	return &AuthorizationController{
		Service: service.AuthorizationServiceHandler(),
	}
}

// AuthorizationControllerInterface ...
type AuthorizationControllerInterface interface {
	PostClaimsClient(context *gin.Context)
	PostClaimsUser(context *gin.Context)
}

// PostClaimsClient ...
func (handler *AuthorizationController) PostClaimsClient(context *gin.Context) {

	result, err := handler.Service.ClientClaimsToken(context)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// PostClaimsUser ...
func (handler *AuthorizationController) PostClaimsUser(context *gin.Context) {

	result, err := handler.Service.UserClaimsToken(context)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}
