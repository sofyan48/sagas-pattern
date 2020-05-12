package entity

import "time"

// Login ...
type Login struct {
	IDUser    string     `gorm:"column:id_user" json:"id_user"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	IDRoles   string     `gorm:"column:id_roles" json:"id_roles"`
	CreatedAt *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
