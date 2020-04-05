package service

import (
	"time"

	"github.com/sofyan48/nemo/src/app/v1/api/user/entity"
	"github.com/sofyan48/nemo/src/app/v1/api/user/event"
)

// UserService ...
type UserService struct {
	Event event.UserEventInterface
}

// UserServiceHandler ...
func UserServiceHandler() *UserService {
	return &UserService{
		Event: event.UserEventHandler(),
	}
}

// UserServiceInterface ...
type UserServiceInterface interface {
	UserCreateService(payload *entity.UserRequest) (*entity.UserResponses, error)
}

// UserCreateService ...
func (service *UserService) UserCreateService(payload *entity.UserRequest) (*entity.UserResponses, error) {
	now := time.Now()
	eventPayload := &entity.UserEvent{}
	eventPayload.Action = "user"
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
