package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sofyan48/svc_user/src/app/v1/entity"
	"github.com/sofyan48/svc_user/src/app/v1/repository"
	"github.com/sofyan48/svc_user/src/utils/logger"
)

// UsersEvent ...
type UsersEvent struct {
	Repository repository.UserRepositoryInterface
	Logger     logger.LoggerInterface
}

// UsersEventHandler ...
func UsersEventHandler() *UsersEvent {
	return &UsersEvent{
		Repository: repository.UserRepositoryHandler(),
		Logger:     logger.LoggerHandler(),
	}
}

// UserEventInterface ...
type UserEventInterface interface {
	InsertDatabase(data *entity.UsersPayload) (*entity.UsersResponse, error)
}

// InsertDatabase ...
func (event *UsersEvent) InsertDatabase(data *entity.UsersPayload) (*entity.UsersResponse, error) {
	now := time.Now()
	userDatabase := &entity.Users{}
	userDatabase.Address = data.Address
	userDatabase.City = data.City
	userDatabase.District = data.District
	userDatabase.Email = data.Email
	userDatabase.FirstName = data.FirstName
	userDatabase.LastName = data.LastName
	userDatabase.PhoneNumber = data.PhoneNumber
	userDatabase.Province = data.Province
	userDatabase.CreatedAt = &now
	userDatabase.UpdatedAt = &now
	err := event.Repository.InsertUsers(userDatabase)
	UUID := uuid.New()
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		event.Logger.Save(UUID.String(), "failed", loggerData)
	}
	loggerData := map[string]interface{}{
		"code":   "400",
		"result": userDatabase,
	}
	_, err = event.Logger.Save(UUID.String(), "succes", loggerData)
	if err != nil {
		return nil, err
	}
	response := &entity.UsersResponse{}
	copier.Copy(&response, &userDatabase)
	return response, nil
}
