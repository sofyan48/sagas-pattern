package routes

import "github.com/gin-gonic/gin"

// ORDERROUTES ...
const ORDERROUTES = VERSION + "/order"

func (rLoader *V2RouterLoader) initOrder(router *gin.Engine) {
	group := router.Group(ORDERROUTES)
	group.POST("", rLoader.Order.OrderCreate)
	group.GET("/get/:uuid", rLoader.Order.GetOrderData)
	group.GET("/list", rLoader.Order.ListOrder)
	group.PUT(":uuid", rLoader.Order.UpdateOrder)
	group.DELETE(":uuid", rLoader.Order.DeleteOrder)
}
