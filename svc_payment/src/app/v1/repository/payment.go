package repository

import (
	"sync"

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
	GetPaymentByID(id int, paymentData *entity.Payment, wg *sync.WaitGroup) error
	GetPaymentList(limit int, offset int) ([]entity.Payment, error)
	InsertPayment(paymentData *entity.Payment, DB *gorm.DB) error
	UpdatePaymentByID(id int, paymentData *entity.Payment, trx *gorm.DB) error
	CheckEmailPayment(email string, paymentData *entity.Payment) bool
}

// GetPaymentByID params
// @id: int
// @paymentData: entity Payment
// wg *sync.WaitGroup
// return error
func (repository *PaymentRepository) GetPaymentByID(id int, paymentData *entity.Payment, wg *sync.WaitGroup) error {
	query := repository.DB.Table("tb_payment")
	query = query.Where("id_payment=?", id)
	query = query.First(&paymentData)
	wg.Done()
	return query.Error
}

// UpdatePaymentByID params
// @id: int
// @paymentData: entity Payment
// return error
func (repository *PaymentRepository) UpdatePaymentByID(id int, paymentData *entity.Payment, trx *gorm.DB) error {
	query := trx.Table("tb_payment")
	query = query.Where("id_payment=?", id)
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
