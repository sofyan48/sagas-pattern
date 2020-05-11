package routes

import "github.com/gin-gonic/gin"

// TOKENROUTES ...
const TOKENROUTES = VERSION + "/token"

func (rLoader *V1RouterLoader) initToken(router *gin.Engine) {
	group := router.Group(TOKENROUTES)
	group.POST("/client", rLoader.Token.PostTokenClient)
}
