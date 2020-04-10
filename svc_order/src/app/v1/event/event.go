package event

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_order/src/app/v1/entity"
	"github.com/sofyan48/svc_order/src/app/v1/repository"
	"github.com/sofyan48/svc_order/src/utils/database"
	"github.com/sofyan48/svc_order/src/utils/logger"
)

// OrderEvent ...
type OrderEvent struct {
	Repository repository.OrderRepositoryInterface
	Logger     logger.LoggerInterface
	DB         *gorm.DB
}

// OrderEventHandler ...
func OrderEventHandler() *OrderEvent {
	return &OrderEvent{
		Repository: repository.OrderRepositoryHandler(),
		Logger:     logger.LoggerHandler(),
		DB:         database.GetTransactionConnection(),
	}
}

// UserEventInterface ...
type UserEventInterface interface {
	InsertDatabase(data *entity.StateFullFormatKafka) (*entity.OrderResponse, error)
	UpdateOrderStatus(data *entity.StateFullFormatKafka) (*entity.OrderResponse, error)
}

// InsertDatabase ...
func (event *OrderEvent) InsertDatabase(data *entity.StateFullFormatKafka) (*entity.OrderResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	orderDatabase := &entity.Order{}
	orderDatabase.UUID = data.UUID
	idOrderType, _ := strconv.ParseInt(data.Data["id_order_type"], 10, 64)
	IDOrderStatus, _ := strconv.ParseInt(data.Data["id_order_status"], 10, 64)
	orderDatabase.IDOrderType = idOrderType
	orderDatabase.IDStatusOrder = IDOrderStatus
	orderDatabase.OrderNumber = data.Data["order_number"]
	orderDatabase.UserUUID = data.Data["uuid_user"]
	orderDatabase.CreatedAt = &now
	orderDatabase.UpdatedAt = &now

	err := event.Repository.InsertOrder(orderDatabase, transaction)
	if err != nil {
		event.DB.Rollback()
		return nil, err
	}
	transaction.Commit()
	response := &entity.OrderResponse{}
	response.UUID = orderDatabase.UUID
	response.OrderNumber = orderDatabase.OrderNumber
	response.IDOrderType = data.Data["id_order_type"]
	response.IDOrderStatus = data.Data["id_order_status"]
	response.UserUUID = data.Data["uuid_user"]
	response.CreatedAt = orderDatabase.CreatedAt
	response.UpdatedAt = orderDatabase.UpdatedAt
	return response, nil
}

// UpdateOrderStatus ...
func (event *OrderEvent) UpdateOrderStatus(data *entity.StateFullFormatKafka) (*entity.OrderResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	orderDatabase := &entity.Order{}
	IDOrderStatus, _ := strconv.ParseInt(data.Data["id_order_status"], 10, 64)
	orderDatabase.IDStatusOrder = IDOrderStatus
	orderDatabase.UpdatedAt = &now
	err := event.Repository.UpdateOrderByUUIID(data.Data["uuid"], orderDatabase, transaction)
	if err != nil {
		event.DB.Rollback()
		return nil, err
	}
	response := &entity.OrderResponse{}
	response.UUID = orderDatabase.UUID
	response.OrderNumber = orderDatabase.OrderNumber
	response.IDOrderType = data.Data["id_order_type"]
	response.IDOrderStatus = data.Data["id_order_status"]
	response.CreatedAt = orderDatabase.CreatedAt
	response.UpdatedAt = orderDatabase.UpdatedAt
	return response, nil
}
