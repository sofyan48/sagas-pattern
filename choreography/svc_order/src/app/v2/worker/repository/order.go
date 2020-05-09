package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_order/src/app/v2/worker/entity"
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
	GetOrderStatus(nmStatus string, orderStatusData *entity.OrderStatus) error
	InsertOrder(usersData *entity.Order, DB *gorm.DB) error
	UpdateOrderByUUIID(uuid string, orderData *entity.Order, trx *gorm.DB) error
}

// GetOrderStatus params
// @id: int
// @orderData: entity Order
// wg *sync.WaitGroup
// return error
func (repository *OrderRepository) GetOrderStatus(nmStatus string, orderStatusData *entity.OrderStatus) error {
	query := repository.DB.Table("tb_order_status")
	query = query.Where("nm_status_order=?", nmStatus)
	query = query.First(&orderStatusData)
	return query.Error
}

// UpdateOrderByUUIID params
// @id: int
// @orderData: entity Order
// return error
func (repository *OrderRepository) UpdateOrderByUUIID(uuid string, orderData *entity.Order, trx *gorm.DB) error {
	query := trx.Table("tb_orders")
	query = query.Where("uuid=?", uuid)
	query = query.Updates(orderData)
	query.Scan(&orderData)
	return query.Error
}

// InsertOrder params
// @orderData: entity Order
// return error
func (repository *OrderRepository) InsertOrder(usersData *entity.Order, DB *gorm.DB) error {
	query := DB.Table("tb_orders")
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}
