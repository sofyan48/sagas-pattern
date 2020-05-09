package entity

import "time"

// PaymentRequest ...
type PaymentRequest struct {
	UUIDOrder         string `json:"uuid_order"`
	UUIDUser          string `json:"uuid_user"`
	IDPaymentStatus   string `json:"id_payment_status"`
	IDPaymentModel    string `json:"id_payment_model"`
	InquiryNumber     string `json:"inquiry_number"`
	BankAccountNumber string `json:"bank_account_number"`
	NMBank            string `json:"nm_bank"`
	PaymentTotal      int    `json:"payment_total"`
}

// PaymentPaidRequest ...
type PaymentPaidRequest struct {
	PaymentTotal      int    `json:"payment_total"`
	PaymentStatus     string `json:"payment_status"`
	BankAccountNumber string `json:"bank_account_number"`
}

// PaymentResponses ...
type PaymentResponses struct {
	UUID      string        `json:"uuid"`
	CreatedAt *time.Time    `json:"created_at"`
	Event     *PaymentEvent `json:"event"`
}
