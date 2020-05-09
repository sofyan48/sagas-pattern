package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_order/src/app/v2/api/order/entity"
	"github.com/sofyan48/svc_order/src/utils/database"
)

// OrderRepository types
type OrderRepository struct {
	DB gorm.DB
}

// OrderRepositoryHandler Order handler repo
// return: OrderRepository
func OrderRepositoryHandler() *OrderRepository {
	return &OrderRepository{
		DB: *database.GetTransactionConnection(),
	}
}

// OrderRepositoryInterface interface
type OrderRepositoryInterface interface {
	GetOrderByID(id int, orderData *entity.Order) error
	GetOrderList(limit int, offset int) ([]entity.Order, error)
	GetOrderByUUID(uuid string, orderStatusData *entity.Order) error
}

// GetOrderByUUID params
// @id: int
// @orderData: entity Order
// wg *sync.WaitGroup
// return error
func (repository *OrderRepository) GetOrderByUUID(uuid string, orderStatusData *entity.Order) error {
	query := repository.DB.Table("tb_orders")
	query = query.Where("uuid=?", uuid)
	query = query.First(&orderStatusData)
	return query.Error
}

// GetOrderByID params
// @id: int
// @orderData: entity Order
// wg *sync.WaitGroup
// return error
func (repository *OrderRepository) GetOrderByID(id int, orderData *entity.Order) error {
	query := repository.DB.Table("tb_orders")
	query = query.Where("id_order=?", id)
	query = query.First(&orderData)
	return query.Error
}

// GetOrderList params
// @id: int
// @orderData: entity Order
// return entity,error
func (repository *OrderRepository) GetOrderList(limit int, offset int) ([]entity.Order, error) {
	users := []entity.Order{}
	query := repository.DB.Table("tb_orders")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}
