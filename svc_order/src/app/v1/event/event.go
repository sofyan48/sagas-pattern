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
}

// InsertDatabase ...
func (event *OrderEvent) InsertDatabase(data *entity.StateFullFormatKafka) (*entity.OrderResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	orderDatabase := &entity.Order{}
	orderDatabase.UUID = data.UUID
	idOrderType, _ := strconv.ParseInt(data.Data["id_order_type"], 10, 64)
	IDPaymentModel, _ := strconv.ParseInt(data.Data["id_payment_model"], 10, 64)
	orderDatabase.IDOrderType = idOrderType
	orderDatabase.IDPaymentModel = IDPaymentModel
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
	response.IDPaymentModel = data.Data["id_payment_model"]
	response.CreatedAt = orderDatabase.CreatedAt
	response.UpdatedAt = orderDatabase.UpdatedAt
	return response, nil
}
