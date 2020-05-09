package controller

import (
	"fmt"
	"time"

	"github.com/sofyan48/svc_payment/src/app/v2/worker/entity"
	"github.com/sofyan48/svc_payment/src/app/v2/worker/event"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
	"github.com/sofyan48/svc_payment/src/utils/logger"
)

// ControllerEvent ....
type ControllerEvent struct {
	Kafka  kafka.KafkaLibraryInterface
	Event  event.PaymentEventInterface
	Logger logger.LoggerInterface
}

// ControllerEventHandler ...
func ControllerEventHandler() *ControllerEvent {
	return &ControllerEvent{
		Kafka:  kafka.KafkaLibraryHandler(),
		Event:  event.PaymentEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// ControllerEventInterface ...
type ControllerEventInterface interface {
	PaymentSave(paymentData *entity.StateFullFormatKafka)
	PaymentPaidOrder(paymentData *entity.StateFullFormatKafka)
}

// PaymentSave ...
func (consumer *ControllerEvent) PaymentSave(paymentData *entity.StateFullFormatKafka) {
	result, err := consumer.Event.InsertDatabase(paymentData)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":   "200",
		"result": result,
	}
	consumer.Logger.Save(paymentData.UUID, "success", loggerData)

	// sending order prepare
	now := time.Now()
	payloadPayment := consumer.Kafka.GetStateFull()
	payloadPayment.Action = "order_update"
	payloadPayment.CreatedAt = &now
	payloadPayment.UUID = result.UUID
	payloadPayment.Data = map[string]interface{}{
		"uuid_order":     result.UUIDOrder,
		"payment_status": "Process",
	}
	resultOrder, _, err := consumer.Kafka.SendEvent("order", payloadPayment)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	paymentLog := map[string]interface{}{
		"code":     "200",
		"messages": "Order Status Update",
		"result":   resultOrder,
	}
	consumer.Logger.Save(paymentData.UUID, "success", paymentLog)

}

// PaymentPaidOrder ...
func (consumer *ControllerEvent) PaymentPaidOrder(paymentData *entity.StateFullFormatKafka) {
	result, err := consumer.Event.PaymentUpdateOrder(paymentData)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		fmt.Println(loggerData)
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":   "200",
		"result": result,
	}
	consumer.Logger.Save(paymentData.UUID, "success", loggerData)

	// sending order prepare
	now := time.Now()
	payloadPayment := consumer.Kafka.GetStateFull()
	payloadPayment.Action = "order_update"
	payloadPayment.CreatedAt = &now
	payloadPayment.UUID = result.UUID
	payloadPayment.Data = map[string]interface{}{
		"uuid_order":     result.UUIDOrder,
		"payment_status": "Waiting",
	}
	resultOrder, _, err := consumer.Kafka.SendEvent("order", payloadPayment)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	paymentLog := map[string]interface{}{
		"code":     "200",
		"messages": "Order Status Update",
		"result":   resultOrder,
	}
	consumer.Logger.Save(paymentData.UUID, "success", paymentLog)
}
