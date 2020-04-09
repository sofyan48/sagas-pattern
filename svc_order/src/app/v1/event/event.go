package event

import (
	"time"

	"github.com/jinzhu/copier"
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

	orderDatabase.CreatedAt = &now
	orderDatabase.UpdatedAt = &now
	err := event.Repository.InsertOrder(orderDatabase, transaction)
	if err != nil {
		event.DB.Rollback()
		return nil, err
	}
	transaction.Commit()
	response := &entity.OrderResponse{}
	copier.Copy(&response, &orderDatabase)
	return response, nil
}
