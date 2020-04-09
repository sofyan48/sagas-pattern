package entity

import "time"

// OrderRequest ...
type OrderRequest struct {
	OrderNumber    string `json:"order_number"`
	UserUUID       string `json:"uuid_user"`
	IDOrderType    string `json:"id_order_type"`
	IDPaymentModel string `json:"id_payment_model"`
}

// OrderResponses ...
type OrderResponses struct {
	UUID      string      `json:"uuid"`
	CreatedAt *time.Time  `json:"created_at"`
	Event     *OrderEvent `json:"event"`
}
