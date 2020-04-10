package entity

import "time"

// OrderPayload Mapping
type OrderPayload struct {
	OrderNumber    string `json:"order_number"`
	UserUUID       string `json:"uuid_user"`
	IDOrderType    string `json:"id_order_type"`
	IDPaymentModel string `json:"id_payment_model"`
}

// OrderResponse Mapping
type OrderResponse struct {
	UUID          string     `json:"uuid"`
	OrderNumber   string     `json:"order_number"`
	UserUUID      string     `json:"uuid_user"`
	IDOrderType   string     `json:"id_order_type"`
	IDOrderStatus string     `json:"id_order_status"`
	CreatedAt     *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// StateFullFormatKafka ...
type StateFullFormatKafka struct {
	UUID      string            `json:"__uuid" bson:"__uuid"`
	Action    string            `json:"__action" bson:"__action"`
	Data      map[string]string `json:"data" bson:"data"`
	CreatedAt *time.Time        `json:"created_at" bson:"created_at"`
}
