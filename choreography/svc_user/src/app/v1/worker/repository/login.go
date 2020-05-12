package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/svc_user/src/app/v1/worker/entity"
)

// UpdateLoginByID params
// @id: int
// @loginData: entity Login
// return error
func (repository *UserRepository) UpdateLoginByID(id int, loginData *entity.Login, trx *gorm.DB) error {
	query := trx.Table("tb_login")
	query = query.Where("id=?", id)
	query = query.Updates(loginData)
	query.Scan(&loginData)
	return query.Error
}

// InsertLogin params
// @loginData: entity Login
// return error
func (repository *UserRepository) InsertLogin(loginData *entity.Login, trx *gorm.DB) error {
	query := trx.Table("tb_login")
	query = query.Create(loginData)
	query.Scan(&loginData)
	return query.Error
}
