package event

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_user/src/app/v1/entity"
	"github.com/sofyan48/svc_user/src/app/v1/repository"
	"github.com/sofyan48/svc_user/src/utils/database"
	"github.com/sofyan48/svc_user/src/utils/logger"
)

// UsersEvent ...
type UsersEvent struct {
	Repository repository.UserRepositoryInterface
	Logger     logger.LoggerInterface
	DB         *gorm.DB
}

// UsersEventHandler ...
func UsersEventHandler() *UsersEvent {
	return &UsersEvent{
		Repository: repository.UserRepositoryHandler(),
		Logger:     logger.LoggerHandler(),
		DB:         database.GetTransactionConnection(),
	}
}

// UserEventInterface ...
type UserEventInterface interface {
	InsertDatabase(data *entity.StateFullFormatKafka) (*entity.UsersResponse, error)
}

// InsertDatabase ...
func (event *UsersEvent) InsertDatabase(data *entity.StateFullFormatKafka) (*entity.UsersResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	userDatabase := &entity.Users{}
	userDatabase.Address = data.Data["address"]
	userDatabase.City = data.Data["city"]
	userDatabase.District = data.Data["district"]
	userDatabase.Email = data.Data["email"]
	userDatabase.FirstName = data.Data["first_name"]
	userDatabase.LastName = data.Data["last_name"]
	userDatabase.PhoneNumber = data.Data["handphone"]
	userDatabase.SiteProfil = data.UUID
	userDatabase.Province = data.Data["province"]
	userDatabase.CreatedAt = &now
	userDatabase.UpdatedAt = &now
	err := event.Repository.InsertUsers(userDatabase, transaction)
	if err != nil {
		event.DB.Rollback()
		return nil, err
	}
	transaction.Commit()
	response := &entity.UsersResponse{}
	copier.Copy(&response, &userDatabase)
	return response, nil
}
