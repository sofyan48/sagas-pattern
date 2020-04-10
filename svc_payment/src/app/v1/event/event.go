package event

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_payment/src/app/v1/entity"
	"github.com/sofyan48/svc_payment/src/app/v1/repository"
	"github.com/sofyan48/svc_payment/src/utils/database"
	"github.com/sofyan48/svc_payment/src/utils/logger"
)

// PaymentEvent ...
type PaymentEvent struct {
	Repository repository.PaymentRepositoryInterface
	Logger     logger.LoggerInterface
	DB         *gorm.DB
}

// PaymentEventHandler ...
func PaymentEventHandler() *PaymentEvent {
	return &PaymentEvent{
		Repository: repository.PaymentRepositoryHandler(),
		Logger:     logger.LoggerHandler(),
		DB:         database.GetTransactionConnection(),
	}
}

// PaymentEventInterface ...
type PaymentEventInterface interface {
	InsertDatabase(data *entity.StateFullFormatKafka) (*entity.PaymentResponse, error)
}

// InsertDatabase ...
func (event *PaymentEvent) InsertDatabase(data *entity.StateFullFormatKafka) (*entity.PaymentResponse, error) {
	transaction := event.DB.Begin()
	now := time.Now()
	paymentDatabase := &entity.Payment{}
	paymentDatabase.UUID = data.UUID
	idPaymentStatus, _ := strconv.ParseInt(data.Data["id_payment_status"], 10, 64)
	IDPaymentModel, _ := strconv.ParseInt(data.Data["id_payment_model"], 10, 64)
	paymentDatabase.IDPaymentStatus = idPaymentStatus
	paymentDatabase.IDPaymentModel = IDPaymentModel
	paymentDatabase.BankAccountNumber = data.Data["bank_account_number"]
	payChange, _ := strconv.Atoi(data.Data["change_total"])
	payTotal, _ := strconv.Atoi(data.Data["payment_total"])
	paymentDatabase.ChangeTotal = payChange
	paymentDatabase.PaymentTotal = payTotal
	dueDate := now.AddDate(0, 0, -1)
	paymentDatabase.DueDate = dueDate
	paymentDatabase.InquiryNumber = data.Data["inquiry_number"]
	paymentDatabase.NMBank = data.Data["nm_bank"]
	paymentDatabase.UUIDOrder = data.Data["uuid_order"]
	paymentDatabase.UUIDUser = data.Data["uuid_user"]
	paymentDatabase.CreatedAt = &now
	paymentDatabase.UpdatedAt = &now

	err := event.Repository.InsertPayment(paymentDatabase, transaction)
	if err != nil {
		event.DB.Rollback()
		return nil, err
	}
	transaction.Commit()
	response := &entity.PaymentResponse{}
	response.UUID = paymentDatabase.UUID
	response.BankAccountNumber = data.Data["bank_account_number"]
	response.ChangeTotal = paymentDatabase.ChangeTotal
	response.DueDate = dueDate
	response.InquiryNumber = paymentDatabase.InquiryNumber
	response.NMBank = paymentDatabase.NMBank
	response.PaymentTotal = payTotal
	response.UUIDUser = paymentDatabase.UUIDUser
	response.UUIDOrder = paymentDatabase.UUIDOrder
	response.IDPaymentModel = data.Data["id_payment_model"]
	response.IDPaymentStatus = data.Data["id_payment_status"]
	response.CreatedAt = paymentDatabase.CreatedAt
	response.UpdatedAt = paymentDatabase.UpdatedAt
	return response, nil
}
