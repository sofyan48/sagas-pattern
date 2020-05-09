package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sofyan48/svc_payment/src/app/v2/swagger/docs"
)

// DOCROUTES ...
const DOCROUTES = VERSION + "/docs"

func (rLoader *V2RouterLoader) initDocs(router *gin.Engine) {
	group := router.Group(DOCROUTES)
	url := ginSwagger.URL("swagger/doc.json")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
