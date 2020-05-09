package event

import (
	"github.com/google/uuid"
	"github.com/sofyan48/svc_payment/src/app/v2/api/payment/entity"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
)

// PAYMENTEVENT ...
const PAYMENTEVENT = "payment"

// PaymentEvent ...
type PaymentEvent struct {
	Kafka kafka.KafkaLibraryInterface
	ready chan bool
}

// PaymentEventHandler ...
func PaymentEventHandler() *PaymentEvent {
	return &PaymentEvent{
		Kafka: kafka.KafkaLibraryHandler(),
		ready: make(chan bool),
	}
}

// PaymentEventInterface ...
type PaymentEventInterface interface {
	PaymentCreateEvent(data *entity.PaymentEvent) (*entity.PaymentEvent, error)
	PaymentCreateEventSequence(data *entity.PaymentEvent) (*entity.PaymentEvent, error)
}

// PaymentCreateEvent ...
func (event *PaymentEvent) PaymentCreateEvent(data *entity.PaymentEvent) (*entity.PaymentEvent, error) {
	format := event.Kafka.GetStateFull()
	format.Action = data.Action
	format.CreatedAt = data.CreatedAt
	format.Data = data.Data
	format.UUID = uuid.New().String()
	data.UUID = format.UUID
	go event.Kafka.SendEvent(PAYMENTEVENT, format)
	data.Status = "QUEUE"
	return data, nil
}

// PaymentCreateEvent ...
func (event *PaymentEvent) PaymentCreateEventSequence(data *entity.PaymentEvent) (*entity.PaymentEvent, error) {
	format := event.Kafka.GetStateFull()
	format.Action = data.Action
	format.CreatedAt = data.CreatedAt
	format.Data = data.Data
	format.UUID = uuid.New().String()
	data.UUID = format.UUID
	event.Kafka.SendEvent(PAYMENTEVENT, format)
	data.Status = "QUEUE"
	return data, nil
}
