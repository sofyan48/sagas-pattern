package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_auth/src/app/v1/api/client/entity"
	"github.com/sofyan48/svc_auth/src/utils/database"
)

// ClientRepository types
type ClientRepository struct {
	DB gorm.DB
}

// ClientRepositoryHandler Clients handler repo
// return: ClientRepository
func ClientRepositoryHandler() *ClientRepository {
	return &ClientRepository{DB: *database.GetTransactionConnection()}
}

// ClientRepositoryInterface interface
type ClientRepositoryInterface interface {
	GetClientByID(id int, clientData *entity.Clients) error
	GetClientsList(limit int, offset int) ([]entity.Clients, error)
	InsertClients(clientData *entity.Clients, DB *gorm.DB) error
	UpdateClientByID(id int, clientData *entity.Clients, trx *gorm.DB) error
	GetClientByClientName(client string) (*entity.Clients, error)
}

// GetClientByID params
// @id: int
// @clientData: entity Clients
// wg *sync.WaitGroup
// return error
func (repository *ClientRepository) GetClientByID(id int, clientData *entity.Clients) error {
	query := repository.DB.Table("tb_client")
	query = query.Where("id=?", id)
	query = query.First(&clientData)
	return query.Error
}

// GetClientByClientName params
// @id: int
// @clientData: entity Clients
// wg *sync.WaitGroup
// return error
func (repository *ClientRepository) GetClientByClientName(client string) (*entity.Clients, error) {
	clientData := &entity.Clients{}
	query := repository.DB.Table("tb_client")
	query = query.Where("client_name=?", client)
	query = query.First(&clientData)
	return clientData, query.Error
}

// UpdateClientByID params
// @id: int
// @clientData: entity Clients
// return error
func (repository *ClientRepository) UpdateClientByID(id int, clientData *entity.Clients, trx *gorm.DB) error {
	query := trx.Table("tb_client")
	query = query.Where("id=?", id)
	query = query.Updates(clientData)
	query.Scan(&clientData)
	return query.Error
}

// GetClientsList params
// @id: int
// @clientData: entity Clients
// return entity,error
func (repository *ClientRepository) GetClientsList(limit int, offset int) ([]entity.Clients, error) {
	users := []entity.Clients{}
	query := repository.DB.Table("tb_client")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}

// InsertClients params
// @clientData: entity Clients
// return error
func (repository *ClientRepository) InsertClients(clientData *entity.Clients, DB *gorm.DB) error {
	query := DB.Table("tb_client")
	query = query.Create(clientData)
	query.Scan(&clientData)
	return query.Error
}
