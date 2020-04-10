package routes

import "github.com/gin-gonic/gin"

// PAYMENTROUTES ...
const PAYMENTROUTES = VERSION + "/payment"

func (rLoader *V1RouterLoader) initPayment(router *gin.Engine) {
	group := router.Group(PAYMENTROUTES)
	group.POST("", rLoader.Payment.PaymentCreate)
	group.GET(":uuid", rLoader.Payment.GetPaymentData)
	group.PUT(":uuid", rLoader.Payment.UpdatePaidPayment)
	group.DELETE(":uuid", rLoader.Payment.DeletePayment)
}
