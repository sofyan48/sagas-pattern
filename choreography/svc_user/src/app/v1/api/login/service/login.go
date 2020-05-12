package service

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/event"
	"github.com/sofyan48/svc_user/src/app/v1/api/login/repository"
	userEntity "github.com/sofyan48/svc_user/src/app/v1/api/user/entity"
	userRepo "github.com/sofyan48/svc_user/src/app/v1/api/user/repository"
	"github.com/sofyan48/svc_user/src/utils/crypto"
	"github.com/sofyan48/svc_user/src/utils/logger"
)

// UserLogin ...
type UserLogin struct {
	Event      event.LoginEventInterface
	Logger     logger.LoggerInterface
	Repository repository.LoginRepositoryInterface
	RepoUser   userRepo.UserRepositoryInterface
	Crypto     crypto.CryptoInterface
}

// UserLoginHandler ...
func UserLoginHandler() *UserLogin {
	return &UserLogin{
		Event:      event.LoginEventHandler(),
		Logger:     logger.LoggerHandler(),
		Repository: repository.LoginRepositoryHandler(),
		RepoUser:   userRepo.UserRepositoryHandler(),
		Crypto:     crypto.CryptoHandler(),
	}
}

// UserLoginInterface ...
type UserLoginInterface interface {
	UserCreateLoginService(payload *entity.UserLoginRequest) (*entity.UserLoginResponses, error)
	GetListLogin(payload *entity.Pagination) (interface{}, error)
	GetLoginByUsername(username string) (interface{}, error)
	CreateSession(payload *entity.GetByUsernameRequest) (interface{}, error)
}

// UserCreateLoginService ...
func (service *UserLogin) UserCreateLoginService(payload *entity.UserLoginRequest) (*entity.UserLoginResponses, error) {
	now := time.Now()
	eventPayload := &entity.UserLoginEvent{}
	eventPayload.Action = "login"
	eventPayload.CreatedAt = &now
	data := map[string]interface{}{
		"id_user":  payload.IDUser,
		"id_roles": payload.IDRoles,
		"username": payload.Username,
		"password": payload.Password,
	}
	eventPayload.Data = data
	event, err := service.Event.LoginCreateEvent(eventPayload)
	if err != nil {
		return nil, err
	}
	result := &entity.UserLoginResponses{}
	result.UUID = event.UUID
	result.Event = event
	result.CreatedAt = event.CreatedAt
	return result, nil
}

// GetListLogin ...
func (service *UserLogin) GetListLogin(payload *entity.Pagination) (interface{}, error) {
	listLogin := []entity.LoginResponse{}
	result, err := service.Repository.GetLoginList(payload.Limit, payload.Page)
	if err != nil {
		return nil, err
	}
	for _, elements := range result {
		dataLogin := entity.LoginResponse{}
		dataLogin.CreatedAt = elements.CreatedAt
		dataLogin.DeletedAt = elements.DeletedAt
		dataLogin.IDRoles = elements.IDRoles
		dataLogin.IDUser = elements.IDUser
		dataLogin.UpdatedAt = elements.UpdatedAt
		dataLogin.Username = elements.Username
		listLogin = append(listLogin, dataLogin)
	}

	return listLogin, nil
}

// GetLoginByUsername ...
func (service *UserLogin) GetLoginByUsername(username string) (interface{}, error) {
	loginData := &entity.Login{}
	err := service.Repository.GetLoginByUsername(username, loginData)
	if err != nil {
		return nil, err
	}
	result := &entity.LoginResponse{}
	copier.Copy(result, loginData)
	return result, nil
}

// CreateSession ...
func (service *UserLogin) CreateSession(payload *entity.GetByUsernameRequest) (interface{}, error) {
	loginData := &entity.Login{}
	err := service.Repository.GetLoginByUsername(payload.Username, loginData)
	if err != nil {
		return nil, err
	}
	if !service.Crypto.CheckPasswordHash(payload.Password, loginData.Password) {
		return nil, errors.New("Password Not Match")
	}
	userData := *&userEntity.Users{}
	err = service.RepoUser.GetUserByID(loginData.IDUser, &userData)
	if err != nil {
		return nil, err
	}
	sessData := &entity.SessionReponse{}
	sessData.Login = loginData
	sessData.Session = userData
	return sessData, nil
}
