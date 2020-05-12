package routes

import "github.com/gin-gonic/gin"

// LOGINROUTES ...
const LOGINROUTES = VERSION + "/login"

func (rLoader *V1RouterLoader) initLogin(router *gin.Engine) {
	group := router.Group(LOGINROUTES)
	group.POST("", rLoader.Login.PostCreateLogin)
	group.GET("/get", rLoader.Middleware.AuthToken(""), rLoader.Login.GetByUsername)
	group.GET("/list", rLoader.Middleware.AuthToken(""), rLoader.Login.GetList)
	// group.PUT(":uuid", rLoader.User.UpdateUser)
	// group.DELETE(":uuid")
}
