package event

import (
	"github.com/google/uuid"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_user/src/utils/kafka"
)

// LOGINEVENT ...
const LOGINEVENT = "users"

// LoginEvent ...
type LoginEvent struct {
	Kafka kafka.KafkaLibraryInterface
}

// LoginEventHandler ...
func LoginEventHandler() *LoginEvent {
	return &LoginEvent{
		Kafka: kafka.KafkaLibraryHandler(),
	}
}

// LoginEventInterface ...
type LoginEventInterface interface {
	LoginCreateEvent(data *entity.UserLoginEvent) (*entity.UserLoginEvent, error)
}

// LoginCreateEvent ...
func (event *LoginEvent) LoginCreateEvent(data *entity.UserLoginEvent) (*entity.UserLoginEvent, error) {
	format := event.Kafka.GetStateFull()
	format.Action = data.Action
	format.CreatedAt = data.CreatedAt
	format.Data = data.Data
	format.UUID = uuid.New().String()
	data.UUID = format.UUID
	go event.Kafka.SendEvent(LOGINEVENT, format)
	data.Status = "QUEUE"
	return data, nil
}
