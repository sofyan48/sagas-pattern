package service

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/entity"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/repository"
	"github.com/sofyan48/svc_auth/src/utils/database"
	"github.com/sofyan48/svc_auth/src/utils/openssl"
)

// ClientService ...
type ClientService struct {
	Repository repository.ClientRepositoryInterface
	SSL        openssl.OpensslInterface
	TrxDB      *gorm.DB
}

// ClientServiceHandler ...
func ClientServiceHandler() *ClientService {
	return &ClientService{
		Repository: repository.ClientRepositoryHandler(),
		SSL:        openssl.OpensslHandler(),
		TrxDB:      database.GetTransactionConnection(),
	}
}

// ClientServiceInterface ...
type ClientServiceInterface interface {
	CreateClient(payload *entity.ClientRequest) (interface{}, error)
}

// CreateClient ...
func (service *ClientService) CreateClient(payload *entity.ClientRequest) (interface{}, error) {
	now := time.Now()
	_, err := service.Repository.GetClientByClientName(payload.ClientName)
	if err == nil {
		return nil, errors.New("Record Exist")
	}

	pathFile, err := service.SSL.GenerateKey(strings.ReplaceAll(strings.ToLower(payload.ClientName), " ", "-"))
	if err != nil {
		return nil, err
	}

	clientData := &entity.Clients{}
	clientData.ClientKey = service.SSL.MD5Hash(payload.ClientName + "-Key")
	clientData.ClientName = payload.ClientName
	clientData.ClientSecret = service.SSL.MD5Hash(payload.ClientName + "-Secret")
	clientData.ClientPUblicKey = pathFile + ".public.pem"
	clientData.ClientPrivateKey = pathFile + ".private.pem"
	clientData.IsFirtsParty = payload.IsFirstParty
	clientData.IsActive = true
	clientData.RedirectUrls = payload.RedirectURIs
	clientData.UpdatedAt = &now
	clientData.CreatedAt = &now
	trx := service.TrxDB.Begin()
	err = service.Repository.InsertClients(clientData, trx)
	if err != nil {
		trx.Rollback()
		return nil, err
	}
	trx.Commit()
	clientDataResponse := &entity.ClientResponses{}
	copier.Copy(clientDataResponse, clientData)
	return clientDataResponse, nil
}
