package routes

import "github.com/gin-gonic/gin"

// PAYMENTROUTES ...
const PAYMENTROUTES = VERSION + "/payment"

func (rLoader *V2RouterLoader) initPayment(router *gin.Engine) {
	group := router.Group(PAYMENTROUTES)
	group.POST("", rLoader.Payment.PaymentCreate)
	group.GET("/get/:uuid", rLoader.Payment.GetPaymentData)
	group.PUT(":uuid", rLoader.Payment.UpdatePaidPayment)
	group.DELETE(":uuid", rLoader.Payment.DeletePayment)
	group.GET("list", rLoader.Payment.ListPayment)
}
