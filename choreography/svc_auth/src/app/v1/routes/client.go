package routes

import "github.com/gin-gonic/gin"

// CLIENTROUTES ...
const CLIENTROUTES = VERSION + "/client"

func (rLoader *V1RouterLoader) initClient(router *gin.Engine) {
	group := router.Group(CLIENTROUTES)
	group.POST("", rLoader.Client.PostClient)
}
