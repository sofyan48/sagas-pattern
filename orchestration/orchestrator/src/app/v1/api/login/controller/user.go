package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/orchestrator/src/app/v1/api/user/entity"
	"github.com/sofyan48/orchestrator/src/app/v1/api/user/service"
	"github.com/sofyan48/orchestrator/src/app/v1/utility/rest"
)

// UserController ...
type UserController struct {
	Service service.UserServiceInterface
}

// UserControllerHandler ...
func UserControllerHandler() *UserController {
	return &UserController{
		Service: service.UserServiceHandler(),
	}
}

// UserControllerInterface ...
type UserControllerInterface interface {
	UserCreate(context *gin.Context)
	GetUserData(context *gin.Context)
}

// UserCreate ...
func (ctrl *UserController) UserCreate(context *gin.Context) {
	payload := &entity.UserRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.UserCreateService(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// GetUserData ...
func (ctrl *UserController) GetUserData(context *gin.Context) {
	uuid := context.Param("uuid")
	result, err := ctrl.Service.UserGetStatus(uuid)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}
