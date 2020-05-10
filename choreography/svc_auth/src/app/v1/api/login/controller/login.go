package controller

import (
	"github.com/sofyan48/svc_auth/src/app/v1/api/Login/service"
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
}
