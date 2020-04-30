package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sofyan48/order/src/app/v1/swagger/docs"
)

// DOCROUTES ...
const DOCROUTES = VERSION + "/docs"

func (rLoader *V1RouterLoader) initDocs(router *gin.Engine) {
	group := router.Group(DOCROUTES)
	url := ginSwagger.URL("swagger/doc.json")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
