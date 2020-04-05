package entity

import "time"

// Users Mapping
type Users struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	FirstName   string     `gorm:"column:first_name;not null;type:varchar(100)" json:"first_name"`
	LastName    string     `gorm:"column:last_name;not null;type:varchar(100)" json:"last_name"`
	Email       string     `gorm:"column:email;unique;not null;type:varchar(100)" json:"email"`
	PhoneNumber string     `gorm:"column:handphone;not null;type:varchar(15)" json:"handphone"`
	Address     string     `gorm:"column:address;type:varchar(255)" json:"address"`
	City        string     `gorm:"column:city;type:varchar(50)" json:"city"`
	Province    string     `gorm:"column:province;type:varchar(50)" json:"province"`
	District    string     `gorm:"column:district;type:varchar(50)" json:"district"`
	CreatedAt   *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
