package service

import (
	"time"

	"github.com/sofyan48/svc_order/src/app/v2/api/order/entity"
	"github.com/sofyan48/svc_order/src/app/v2/api/order/event"
	"github.com/sofyan48/svc_order/src/app/v2/api/order/repository"
)

// OrderService ...
type OrderService struct {
	Event      event.OrderEventInterface
	Repository repository.OrderRepositoryInterface
}

// OrderServiceHandler ...
func OrderServiceHandler() *OrderService {
	return &OrderService{
		Event:      event.OrderEventHandler(),
		Repository: repository.OrderRepositoryHandler(),
	}
}

// OrderServiceInterface ...
type OrderServiceInterface interface {
	OrderCreateService(payload *entity.OrderRequest) (*entity.OrderResponses, error)
	OrderGetUUID(uuid string) (interface{}, error)
	OrderList(pagination *entity.Pagination) (interface{}, error)
}

// OrderCreateService ...
func (service *OrderService) OrderCreateService(payload *entity.OrderRequest) (*entity.OrderResponses, error) {
	now := time.Now()
	eventPayload := &entity.OrderEvent{}
	eventPayload.Action = "order_save"
	eventPayload.CreatedAt = &now
	data := map[string]interface{}{
		"order_number":      payload.OrderNumber,
		"uuid_user":         payload.UserUUID,
		"id_order_type":     payload.IDOrderType,
		"id_order_status":   payload.IDOrderStatus,
		"id_payment_status": payload.IDPaymentStatus,
		"id_payment_model":  payload.IDPaymentModel,
		"inquiry_number":    payload.InquiryNumber,
		"payment_order":     payload.PaymentOrder,
		"nm_bank":           payload.NMBank,
	}
	eventPayload.Data = data
	event, err := service.Event.OrderCreateEvent(eventPayload)
	if err != nil {
		return nil, err
	}
	result := &entity.OrderResponses{}
	result.UUID = event.UUID
	result.Event = event
	result.CreatedAt = event.CreatedAt
	return result, nil
}

// OrderGetUUID ...
func (service *OrderService) OrderGetUUID(uuid string) (interface{}, error) {
	orderData := &entity.Order{}
	err := service.Repository.GetOrderByUUID(uuid, orderData)
	if err != nil {
		return nil, err
	}
	return orderData, nil
}

// OrderList ...
func (service *OrderService) OrderList(pagination *entity.Pagination) (interface{}, error) {
	listOrder := []entity.Order{}
	listOrder, err := service.Repository.GetOrderList(pagination.Limit, pagination.Page)
	if err != nil {
		return nil, err
	}
	return listOrder, nil
}
