package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_order/src/app/v2/api/order/entity"
	"github.com/sofyan48/svc_order/src/app/v2/api/order/service"
	"github.com/sofyan48/svc_order/src/app/v2/utility/rest"
)

// OrderController ...
type OrderController struct {
	Service service.OrderServiceInterface
}

// OrderControllerHandler ...
func OrderControllerHandler() *OrderController {
	return &OrderController{
		Service: service.OrderServiceHandler(),
	}
}

// OrderControllerInterface ...
type OrderControllerInterface interface {
	OrderCreate(context *gin.Context)
	UpdateOrder(context *gin.Context)
	GetOrderData(context *gin.Context)
	DeleteOrder(context *gin.Context)
	ListOrder(context *gin.Context)
}

// OrderCreate ...
func (ctrl *OrderController) OrderCreate(context *gin.Context) {
	payload := &entity.OrderRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.OrderCreateService(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// GetOrderData ...
func (ctrl *OrderController) GetOrderData(context *gin.Context) {
	uuid := context.Param("uuid")
	result, err := ctrl.Service.OrderGetUUID(uuid)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// ListOrder ...
func (ctrl *OrderController) ListOrder(context *gin.Context) {
	pagination := &entity.Pagination{}
	context.ShouldBind(pagination)
	result, err := ctrl.Service.OrderList(pagination)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseList(context, http.StatusOK, result, pagination)
	return
}

// UpdateOrder ...
func (ctrl *OrderController) UpdateOrder(context *gin.Context) {
	rest.ResponseMessages(context, http.StatusOK, "OK")
}

// DeleteOrder ...
func (ctrl *OrderController) DeleteOrder(context *gin.Context) {
	rest.ResponseMessages(context, http.StatusOK, "OK")
}
