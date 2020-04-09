package repository

import (
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_order/src/app/v1/entity"
	"github.com/sofyan48/svc_order/src/utils/database"
)

// OrderRepository types
type OrderRepository struct {
	DB gorm.DB
}

// OrderRepositoryHandler Order handler repo
// return: OrderRepository
func OrderRepositoryHandler() *OrderRepository {
	return &OrderRepository{DB: *database.GetTransactionConnection()}
}

// OrderRepositoryInterface interface
type OrderRepositoryInterface interface {
	GetUserByID(id int, userData *entity.Order, wg *sync.WaitGroup) error
	GetOrderList(limit int, offset int) ([]entity.Order, error)
	InsertOrder(usersData *entity.Order, DB *gorm.DB) error
	UpdateUserByID(id int, userData *entity.Order, trx *gorm.DB) error
	CheckEmailOrder(email string, usersData *entity.Order) bool
}

// GetUserByID params
// @id: int
// @userData: entity Order
// wg *sync.WaitGroup
// return error
func (repository *OrderRepository) GetUserByID(id int, userData *entity.Order, wg *sync.WaitGroup) error {
	query := repository.DB.Table("tb_orders")
	query = query.Where("id_order=?", id)
	query = query.First(&userData)
	wg.Done()
	return query.Error
}

// UpdateUserByID params
// @id: int
// @userData: entity Order
// return error
func (repository *OrderRepository) UpdateUserByID(id int, userData *entity.Order, trx *gorm.DB) error {
	query := trx.Table("tb_orders")
	query = query.Where("id_order=?", id)
	query = query.Updates(userData)
	query.Scan(&userData)
	return query.Error
}

// GetOrderList params
// @id: int
// @userData: entity Order
// return entity,error
func (repository *OrderRepository) GetOrderList(limit int, offset int) ([]entity.Order, error) {
	users := []entity.Order{}
	query := repository.DB.Table("tb_orders")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}

// InsertOrder params
// @userData: entity Order
// return error
func (repository *OrderRepository) InsertOrder(usersData *entity.Order, DB *gorm.DB) error {
	query := DB.Table("tb_orders")
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}

// CheckEmailOrder params
// @email : string
// @userData: entity Order
// return error
func (repository *OrderRepository) CheckEmailOrder(email string, usersData *entity.Order) bool {
	query := repository.DB.Table("tb_orders")
	if err := query.Where("email=?", email).First(&usersData).Error; err != nil {
		return false
	}
	return true
}
