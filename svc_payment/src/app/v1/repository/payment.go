package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_payment/src/app/v1/entity"
	"github.com/sofyan48/svc_payment/src/utils/database"
)

// PaymentRepository types
type PaymentRepository struct {
	DB gorm.DB
}

// PaymentRepositoryHandler Payment handler repo
// return: PaymentRepository
func PaymentRepositoryHandler() *PaymentRepository {
	return &PaymentRepository{DB: *database.GetTransactionConnection()}
}

// PaymentRepositoryInterface interface
type PaymentRepositoryInterface interface {
	GetPaymentByOrder(uuidOrder string, paymentData *entity.Payment) error
	GetPaymentList(limit int, offset int) ([]entity.Payment, error)
	GetPaymentStatus(status string, paymentStatus *entity.PaymentStatus) error
	InsertPayment(paymentData *entity.Payment, DB *gorm.DB) error
	UpdatePaymentByOrder(uuidOrder string, paymentData *entity.Payment, trx *gorm.DB) error
	CheckEmailPayment(email string, paymentData *entity.Payment) bool
}

// GetPaymentByOrder params
// @id: int
// @paymentData: entity Payment
// wg *sync.WaitGroup
// return error
func (repository *PaymentRepository) GetPaymentByOrder(uuidOrder string, paymentData *entity.Payment) error {
	query := repository.DB.Table("tb_payment")
	query = query.Where("uuid_order=?", uuidOrder)
	query = query.First(&paymentData)
	return query.Error
}

// GetPaymentStatus ...
func (repository *PaymentRepository) GetPaymentStatus(status string, paymentStatus *entity.PaymentStatus) error {
	query := repository.DB.Table("tb_payment_status")
	query = query.Where("nm_payment_status=?", status)
	query = query.First(&paymentStatus)
	return query.Error
}

// UpdatePaymentByOrder params
// @id: int
// @paymentData: entity Payment
// return error
func (repository *PaymentRepository) UpdatePaymentByOrder(uuidOrder string, paymentData *entity.Payment, trx *gorm.DB) error {
	query := trx.Table("tb_payment")
	query = query.Where("uuid_order=?", uuidOrder)
	query = query.Updates(paymentData)
	query.Scan(&paymentData)
	return query.Error
}

// GetPaymentList params
// @id: int
// @paymentData: entity Payment
// return entity,error
func (repository *PaymentRepository) GetPaymentList(limit int, offset int) ([]entity.Payment, error) {
	payments := []entity.Payment{}
	query := repository.DB.Table("tb_payment")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&payments)
	return payments, query.Error
}

// InsertPayment params
// @paymentData: entity Payment
// return error
func (repository *PaymentRepository) InsertPayment(paymentData *entity.Payment, DB *gorm.DB) error {
	query := DB.Table("tb_payment")
	query = query.Create(paymentData)
	query.Scan(&paymentData)
	return query.Error
}

// CheckEmailPayment params
// @email : string
// @paymentData: entity Payment
// return error
func (repository *PaymentRepository) CheckEmailPayment(email string, paymentData *entity.Payment) bool {
	query := repository.DB.Table("tb_payment")
	if err := query.Where("email=?", email).First(&paymentData).Error; err != nil {
		return false
	}
	return true
}
