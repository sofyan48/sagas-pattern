package routes

import (
	"github.com/gin-gonic/gin"
)

// HEALTHROUTES ...
const HEALTHROUTES = VERSION + "/health"

func (rLoader *V1RouterLoader) initHealth(router *gin.Engine) {
	group := router.Group(HEALTHROUTES)
	group.GET("", rLoader.Health.Health)
}
