package routes

import "github.com/gin-gonic/gin"

// CLIENTROUTES ...
const LOGINROUTES = VERSION + "/login"

func (rLoader *V1RouterLoader) initLogin(router *gin.Engine) {
	group := router.Group(LOGINROUTES)
	group.POST("", rLoader.Middleware.AuthToken(), rLoader.Login.PostLoginClient)
}
