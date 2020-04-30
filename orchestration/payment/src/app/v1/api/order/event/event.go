package event

import (
	"github.com/google/uuid"
	"github.com/sofyan48/payment/src/app/v1/api/order/entity"
	"github.com/sofyan48/payment/src/app/v1/utility/kafka"
)

// ORDEREVENT ...
const ORDEREVENT = "order"

// OrderEvent ...
type OrderEvent struct {
	Kafka kafka.KafkaLibraryInterface
}

// OrderEventHandler ...
func OrderEventHandler() *OrderEvent {
	return &OrderEvent{
		Kafka: kafka.KafkaLibraryHandler(),
	}
}

// OrderEventInterface ...
type OrderEventInterface interface {
	OrderCreateEvent(data *entity.OrderEvent) (*entity.OrderEvent, error)
}

// OrderCreateEvent ...
func (event *OrderEvent) OrderCreateEvent(data *entity.OrderEvent) (*entity.OrderEvent, error) {
	format := event.Kafka.GetStateFull()
	format.Action = data.Action
	format.CreatedAt = data.CreatedAt
	format.Data = data.Data
	format.UUID = uuid.New().String()
	data.UUID = format.UUID
	go event.Kafka.SendEvent(ORDEREVENT, format)
	data.Status = "QUEUE"
	return data, nil
}
