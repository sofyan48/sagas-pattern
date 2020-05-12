package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_user/src/app/v1/worker/entity"
	"github.com/sofyan48/svc_user/src/utils/database"
)

// UserRepository types
type UserRepository struct {
	DB gorm.DB
}

// UserRepositoryHandler Users handler repo
// return: UserRepository
func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *database.GetTransactionConnection()}
}

// UserRepositoryInterface interface
type UserRepositoryInterface interface {
	InsertUsers(usersData *entity.Users, DB *gorm.DB) error
	UpdateUserByID(id int, userData *entity.Users, trx *gorm.DB) error

	InsertLogin(loginData *entity.Login, trx *gorm.DB) error
	UpdateLoginByID(id int, loginData *entity.Login, trx *gorm.DB) error
}

// UpdateUserByID params
// @id: int
// @userData: entity Users
// return error
func (repository *UserRepository) UpdateUserByID(id int, userData *entity.Users, trx *gorm.DB) error {
	query := trx.Table("tb_users")
	query = query.Where("id_user=?", id)
	query = query.Updates(userData)
	query.Scan(&userData)
	return query.Error
}

// InsertUsers params
// @userData: entity Users
// return error
func (repository *UserRepository) InsertUsers(usersData *entity.Users, DB *gorm.DB) error {
	query := DB.Table("tb_users")
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}
