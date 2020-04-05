package routes

import "github.com/gin-gonic/gin"

// USERROUTES ...
const USERROUTES = VERSION + "/user"

func (rLoader *V1RouterLoader) initUser(router *gin.Engine) {
	group := router.Group(USERROUTES)
	group.POST("", rLoader.User.UserCreate)
}
