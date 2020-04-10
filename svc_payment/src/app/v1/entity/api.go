package entity

import "time"

// PaymentPayload Mapping
type PaymentPayload struct {
	UUIDOrder         string `json:"uuid_order"`
	UUIDUser          string `json:"uuid_user"`
	IDPaymentStatus   string `json:"id_payment_status"`
	IDPaymentModel    string `json:"id_payment_model"`
	InquiryNumber     string `json:"inquiry_number"`
	BankAccountNumber string `json:"bank_account_number"`
	NMBank            string `json:"nm_bank"`
	PaymentTotal      int    `json:"payment_total"`
	ChangeTotal       int    `json:"change_total"`
}

// PaymentResponse Mapping
type PaymentResponse struct {
	UUID              string     `json:"uuid"`
	UUIDOrder         string     `json:"uuid_order"`
	UUIDUser          string     `json:"uuid_user"`
	IDPaymentStatus   string     `json:"id_payment_status"`
	IDPaymentModel    string     `json:"id_payment_model"`
	InquiryNumber     string     `json:"inquiry_number"`
	BankAccountNumber string     `json:"bank_account_number"`
	NMBank            string     `json:"nm_bank"`
	PaymentTotal      int        `json:"payment_total"`
	ChangeTotal       int        `json:"change_total"`
	DueDate           time.Time  `json:"due_date"`
	CreatedAt         *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// StateFullFormatKafka ...
type StateFullFormatKafka struct {
	UUID      string            `json:"__uuid" bson:"__uuid"`
	Action    string            `json:"__action" bson:"__action"`
	Data      map[string]string `json:"data" bson:"data"`
	CreatedAt *time.Time        `json:"created_at" bson:"created_at"`
}
