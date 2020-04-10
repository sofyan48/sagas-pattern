package controller

import (
	"time"

	"github.com/sofyan48/svc_order/src/app/v1/entity"
	"github.com/sofyan48/svc_order/src/app/v1/event"
	"github.com/sofyan48/svc_order/src/utils/kafka"
	"github.com/sofyan48/svc_order/src/utils/logger"
)

// ControllerEvent ....
type ControllerEvent struct {
	Kafka  kafka.KafkaLibraryInterface
	Event  event.UserEventInterface
	Logger logger.LoggerInterface
}

// ControllerEventHandler ...
func ControllerEventHandler() *ControllerEvent {
	return &ControllerEvent{
		Kafka:  kafka.KafkaLibraryHandler(),
		Event:  event.OrderEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// ControllerEventInterface ...
type ControllerEventInterface interface {
	OrderLoad(dataOrder *entity.StateFullFormatKafka)
	UpdateOrder(dataOrder *entity.StateFullFormatKafka)
}

// OrderLoad ...
func (consumer *ControllerEvent) OrderLoad(dataOrder *entity.StateFullFormatKafka) {
	result, err := consumer.Event.InsertDatabase(dataOrder)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(dataOrder.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":     "200",
		"messages": "Insert Order Success",
		"result":   result,
	}
	consumer.Logger.Save(dataOrder.UUID, "success", loggerData)

	// sending payment prepare
	now := time.Now()
	payloadPayment := consumer.Kafka.GetStateFull()
	payloadPayment.Action = "payment_save"
	payloadPayment.CreatedAt = &now
	payloadPayment.UUID = result.UUID
	payloadPayment.Data = map[string]interface{}{
		"uuid_order":        result.UUID,
		"uuid_user":         result.UserUUID,
		"id_payment_status": dataOrder.Data["id_payment_status"],
		"id_payment_model":  dataOrder.Data["id_payment_model"],
		"payment_order":     dataOrder.Data["payment_order"],
		"inquiry_number":    dataOrder.Data["inquiry_number"],
		"nm_bank":           dataOrder.Data["nm_bank"],
	}
	resultPayment, _, err := consumer.Kafka.SendEvent("payment", payloadPayment)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(dataOrder.UUID, "failed", loggerData)
		return
	}
	paymentLog := map[string]interface{}{
		"code":     "200",
		"messages": "Payment Created | Offset",
		"result":   resultPayment,
	}
	consumer.Logger.Save(dataOrder.UUID, "success", paymentLog)
}

// updateOrder ...
func (consumer *ControllerEvent) UpdateOrder(dataOrder *entity.StateFullFormatKafka) {
	result, err := consumer.Event.UpdateOrderStatus(dataOrder)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(dataOrder.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":     "200",
		"messages": "Order Update Success",
		"result":   result,
	}
	consumer.Logger.Save(dataOrder.UUID, "success", loggerData)
}
