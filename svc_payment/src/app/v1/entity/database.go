package entity

import "time"

// Payment Mapping
type Payment struct {
	ID                int        `gorm:"column:id_payment;primary_key" json:"id_payment"`
	UUID              string     `gorm:"column:uuid;primary_key" json:"uuid"`
	UUIDOrder         string     `gorm:"column:uuid_order;primary_key" json:"uuid_order"`
	UUIDUser          string     `gorm:"column:uuid_user;primary_key" json:"uuid_user"`
	IDPaymentStatus   int64      `gorm:"column:id_payment_status;not null;type:int(100)" json:"id_payment_status"`
	IDPaymentModel    int64      `gorm:"column:id_payment_model;not null;type:int(100)" json:"id_payment_model"`
	InquiryNumber     string     `gorm:"column:inquiry_number;not null;type:varchar(100)" json:"inquiry_number"`
	BankAccountNumber string     `gorm:"column:bank_account_number;not null;type:varchar(100)" json:"bank_account_number"`
	NMBank            string     `gorm:"column:nm_bank;not null;type:varchar(100)" json:"nm_bank"`
	PaymentTotal      int        `gorm:"column:payment_total;not null;type:varchar(100)" json:"payment_total"`
	PaymentOrder      int        `gorm:"column:payment_order;not null;type:int(100)" json:"payment_order"`
	ChangeTotal       int        `gorm:"column:change_total;not null;type:varchar(100)" json:"change_total"`
	DueDate           time.Time  `gorm:"column:due_date" json:"due_date"`
	CreatedAt         *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
