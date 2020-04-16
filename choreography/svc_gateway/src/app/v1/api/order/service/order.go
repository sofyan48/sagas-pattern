package service

import (
	"time"

	"github.com/sofyan48/svc_gateway/src/app/v1/api/order/entity"
	"github.com/sofyan48/svc_gateway/src/app/v1/api/order/event"
	"github.com/sofyan48/svc_gateway/src/app/v1/utility/logger"
)

// OrderService ...
type OrderService struct {
	Event  event.OrderEventInterface
	Logger logger.LoggerInterface
}

// OrderServiceHandler ...
func OrderServiceHandler() *OrderService {
	return &OrderService{
		Event:  event.OrderEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// OrderServiceInterface ...
type OrderServiceInterface interface {
	OrderCreateService(payload *entity.OrderRequest) (*entity.OrderResponses, error)
	OrderGetStatus(uuid string) (interface{}, error)
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

// OrderGetStatus ...
func (service *OrderService) OrderGetStatus(uuid string) (interface{}, error) {
	data, err := service.Logger.Find(uuid, "order")
	if err != nil {
		return nil, err
	}
	return data, nil
}
