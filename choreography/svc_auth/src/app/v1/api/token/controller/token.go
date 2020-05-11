package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/token/entity"
	"github.com/sofyan48/svc_auth/src/app/v1/api/token/service"
	"github.com/sofyan48/svc_auth/src/utils/rest"
)

// TokenController ...
type TokenController struct {
	Service service.TokenServiceInterface
}

// TokenControllerHandler ...
func TokenControllerHandler() *TokenController {
	return &TokenController{
		Service: service.TokenServiceHandler(),
	}
}

// TokenControllerInterface ...
type TokenControllerInterface interface {
	PostTokenClient(context *gin.Context)
}

// PostTokenClient ...
func (handler *TokenController) PostTokenClient(context *gin.Context) {
	payload := &entity.ClientCredentialRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	result, err := handler.Service.ClientCredential(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
}
