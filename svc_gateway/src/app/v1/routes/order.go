package routes

import "github.com/gin-gonic/gin"

// ORDERROUTES ...
const ORDERROUTES = VERSION + "/order"

func (rLoader *V1RouterLoader) initOrder(router *gin.Engine) {
	group := router.Group(USERROUTES)
	group.POST("", rLoader.Order.OrderCreate)
	group.GET(":uuid", rLoader.Order.GetOrderData)
	group.PUT(":uuid", rLoader.Order.UpdateOrder)
	group.DELETE(":uuid", rLoader.Order.DeleteOrder)
}
