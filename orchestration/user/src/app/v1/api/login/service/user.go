package service

import (
	"fmt"
	"time"

	"github.com/sofyan48/user/src/app/v1/api/user/entity"
	"github.com/sofyan48/user/src/app/v1/api/user/event"
	"github.com/sofyan48/user/src/app/v1/utility/logger"
)

// UserService ...
type UserService struct {
	Event  event.UserEventInterface
	Logger logger.LoggerInterface
}

// UserServiceHandler ...
func UserServiceHandler() *UserService {
	return &UserService{
		Event:  event.UserEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// UserServiceInterface ...
type UserServiceInterface interface {
	UserCreateService(payload *entity.UserRequest) (*entity.UserResponses, error)
	UserGetStatus(uuid string) (interface{}, error)
}

// UserCreateService ...
func (service *UserService) UserCreateService(payload *entity.UserRequest) (*entity.UserResponses, error) {
	now := time.Now()
	eventPayload := &entity.UserEvent{}
	eventPayload.Action = "users"
	eventPayload.CreatedAt = &now
	data := map[string]interface{}{
		"first_name": payload.FirstName,
		"last_name":  payload.LastName,
		"email":      payload.Email,
		"handphone":  payload.PhoneNumber,
		"address":    payload.Address,
		"city":       payload.City,
		"province":   payload.Province,
		"district":   payload.District,
	}
	eventPayload.Data = data
	event, err := service.Event.UserCreateEvent(eventPayload)
	if err != nil {
		return nil, err
	}
	result := &entity.UserResponses{}
	result.UUID = event.UUID
	result.Event = event
	result.CreatedAt = event.CreatedAt
	return result, nil
}

// UserGetStatus ...
func (service *UserService) UserGetStatus(uuid string) (interface{}, error) {
	fmt.Println("UUID", uuid)
	data, err := service.Logger.Find(uuid, "users")
	if err != nil {
		return nil, err
	}
	return data, nil
}
