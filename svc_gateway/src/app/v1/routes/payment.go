package routes

import "github.com/gin-gonic/gin"

// PAYMENTROUTES ...
const PAYMENTROUTES = VERSION + "/payment"

func (rLoader *V1RouterLoader) initPayment(router *gin.Engine) {
	group := router.Group(ORDERROUTES)
	group.POST("", rLoader.Payment.PaymentCreate)
	group.GET(":uuid", rLoader.Payment.GetPaymentData)
	group.PUT(":uuid", rLoader.Payment.UpdatePayment)
	group.DELETE(":uuid", rLoader.Payment.DeletePayment)
}
