package routes

import "github.com/gin-gonic/gin"

// AUTHROUTES ...
const AUTHROUTES = VERSION + "/authorization"

func (rLoader *V1RouterLoader) initAuthService(router *gin.Engine) {
	group := router.Group(AUTHROUTES)
	group.GET("/client", rLoader.Authorizer.PostClaimsClient)
	group.GET("/user", rLoader.Authorizer.PostClaimsUser)
}
