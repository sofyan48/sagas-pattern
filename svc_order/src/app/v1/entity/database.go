package entity

import "time"

// Order Mapping
type Order struct {
	ID             uint64     `gorm:"column:id_order;primary_key" json:"id_order"`
	UUID           string     `gorm:"column:uuid;primary_key" json:"uuid"`
	OrderNumber    string     `gorm:"column:order_number;not null;type:varchar(100)" json:"order_number"`
	UserUUID       string     `gorm:"column:uuid_user;not null;type:varchar(100)" json:"uuid_user"`
	IDOrderType    string     `gorm:"column:id_order_type;unique;not null;type:varchar(100)" json:"id_order_type"`
	IDPaymentModel string     `gorm:"column:id_payment_model;unique;not null;type:varchar(100)" json:"id_payment_model"`
	PhoneNumber    string     `gorm:"column:handphone;not null;type:varchar(15)" json:"handphone"`
	CreatedAt      *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
