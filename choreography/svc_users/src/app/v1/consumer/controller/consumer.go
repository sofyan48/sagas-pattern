package controller

import (
	"fmt"

	"github.com/sofyan48/svc_user/src/app/v1/entity"
	"github.com/sofyan48/svc_user/src/app/v1/event"
	"github.com/sofyan48/svc_user/src/utils/kafka"
	"github.com/sofyan48/svc_user/src/utils/logger"
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
		Event:  event.UsersEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// ControllerEventInterface ...
type ControllerEventInterface interface {
	UserLoad(dataUser *entity.StateFullFormatKafka)
}

// UserLoad ...
func (consumer *ControllerEvent) UserLoad(dataUser *entity.StateFullFormatKafka) {
	result, err := consumer.Event.InsertDatabase(dataUser)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(dataUser.UUID, "failed", loggerData)
	}
	loggerData := map[string]interface{}{
		"code":   "200",
		"result": result,
	}
	data, err := consumer.Logger.Save(dataUser.UUID, "success", loggerData)
	fmt.Println(data, err)
}
