package event

import (
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_user/src/app/v1/worker/entity"
	"github.com/sofyan48/svc_user/src/app/v1/worker/repository"
	"github.com/sofyan48/svc_user/src/utils/crypto"
	"github.com/sofyan48/svc_user/src/utils/database"
	"github.com/sofyan48/svc_user/src/utils/logger"
)

// UsersEvent ...
type UsersEvent struct {
	Repository repository.UserRepositoryInterface
	Logger     logger.LoggerInterface
	DB         *gorm.DB
	Crypto     crypto.CryptoInterface
}

// UsersEventHandler ...
func UsersEventHandler() *UsersEvent {
	return &UsersEvent{
		Repository: repository.UserRepositoryHandler(),
		Logger:     logger.LoggerHandler(),
		DB:         database.GetTransactionConnection(),
		Crypto:     crypto.CryptoHandler(),
	}
}

// UserEventInterface ...
type UserEventInterface interface {
	InsertDatabase(data *entity.StateFullFormatKafka) (*entity.UsersResponse, error)
	InserLogin(data *entity.StateFullFormatKafka) (*entity.LoginResponse, error)
}

// InsertDatabase ...
func (event *UsersEvent) InsertDatabase(data *entity.StateFullFormatKafka) (*entity.UsersResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	userDatabase := &entity.Users{}
	userDatabase.Address = data.Data["address"]
	userDatabase.UUID = data.UUID
	userDatabase.City = data.Data["city"]
	userDatabase.District = data.Data["district"]
	userDatabase.Email = data.Data["email"]
	userDatabase.FirstName = data.Data["first_name"]
	userDatabase.LastName = data.Data["last_name"]
	userDatabase.PhoneNumber = data.Data["handphone"]
	userDatabase.SiteProfil = strings.SplitAfter(userDatabase.FirstName, " ")[0] + "-" + data.UUID
	userDatabase.Province = data.Data["province"]
	userDatabase.CreatedAt = &now
	userDatabase.UpdatedAt = &now
	err := event.Repository.InsertUsers(userDatabase, transaction)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	response := &entity.UsersResponse{}
	copier.Copy(&response, &userDatabase)
	return response, nil
}

// InserLogin ...
func (event *UsersEvent) InserLogin(data *entity.StateFullFormatKafka) (*entity.LoginResponse, error) {
	trx := event.DB.Begin()
	now := time.Now()
	loginDatabase := &entity.Login{}
	loginDatabase.IDRoles = data.Data["id_roles"]
	loginDatabase.IDUser = data.Data["id_user"]
	loginDatabase.Password, _ = event.Crypto.HashPassword(data.Data["password"])
	loginDatabase.Username = data.Data["username"]
	loginDatabase.CreatedAt = &now
	loginDatabase.UpdatedAt = &now

	err := event.Repository.InsertLogin(loginDatabase, trx)
	if err != nil {
		trx.Rollback()
		return nil, err
	}
	trx.Commit()
	response := &entity.LoginResponse{}
	copier.Copy(&response, &loginDatabase)
	return response, nil
}
