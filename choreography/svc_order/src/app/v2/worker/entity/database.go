package entity

import "time"

// Order Mapping
type Order struct {
	ID            int        `gorm:"column:id_order;primary_key" json:"id_order"`
	UUID          string     `gorm:"column:uuid:not null" json:"uuid"`
	OrderNumber   string     `gorm:"column:order_number;not null;type:varchar(100)" json:"order_number"`
	UserUUID      string     `gorm:"column:user_uuid;not null;type:varchar(100)" json:"uuid_user"`
	IDOrderType   int64      `gorm:"column:id_order_type;unique;not null;type:varchar(100)" json:"id_order_type"`
	IDStatusOrder int64      `gorm:"column:id_order_status;unique;not null;type:varchar(100)" json:"id_order_status"`
	CreatedAt     *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// OrderStatus Mapping
type OrderStatus struct {
	ID            int64      `gorm:"column:id_order_status;primary_key" json:"id_order_status"`
	NMStatusOrder string     `gorm:"column:nm_status_order;primary_key" json:"nm_status_order"`
	CreatedAt     *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
