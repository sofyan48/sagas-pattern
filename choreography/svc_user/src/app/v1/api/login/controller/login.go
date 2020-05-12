package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/service"
	"github.com/sofyan48/svc_user/src/utils/rest"
)

// LoginController ...
type LoginController struct {
	Service service.UserLoginInterface
}

// LoginControllerHandler ...
func LoginControllerHandler() *LoginController {
	return &LoginController{
		Service: service.UserLoginHandler(),
	}
}

// LoginControllerInterface ...
type LoginControllerInterface interface {
	PostCreateLogin(context *gin.Context)
	GetList(context *gin.Context)
	GetByUsername(context *gin.Context)
	PostSessionData(context *gin.Context)
}

// PostCreateLogin ...
func (handler *LoginController) PostCreateLogin(context *gin.Context) {
	payload := &entity.UserLoginRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	result, err := handler.Service.UserCreateLoginService(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// GetList ...
func (handler *LoginController) GetList(context *gin.Context) {
	payload := &entity.Pagination{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	result, err := handler.Service.GetListLogin(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseList(context, http.StatusOK, result, payload)
	return
}

// GetByUsername ...
func (handler *LoginController) GetByUsername(context *gin.Context) {
	payload := &entity.GetByUsernameRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	result, err := handler.Service.GetLoginByUsername(payload.Username)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// PostSessionData ...
func (handler *LoginController) PostSessionData(context *gin.Context) {
	payload := &entity.GetByUsernameRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	result, err := handler.Service.CreateSession(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}
