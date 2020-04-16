package repository

import (
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_user/src/app/v1/entity"
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
	GetUserByID(id int, userData *entity.Users, wg *sync.WaitGroup) error
	GetUsersList(limit int, offset int) ([]entity.Users, error)
	InsertUsers(usersData *entity.Users, DB *gorm.DB) error
	UpdateUserByID(id int, userData *entity.Users, trx *gorm.DB) error
	CheckEmailUsers(email string, usersData *entity.Users) bool
}

// GetUserByID params
// @id: int
// @userData: entity Users
// wg *sync.WaitGroup
// return error
func (repository *UserRepository) GetUserByID(id int, userData *entity.Users, wg *sync.WaitGroup) error {
	query := repository.DB.Table("tb_users")
	query = query.Where("id_user=?", id)
	query = query.First(&userData)
	wg.Done()
	return query.Error
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

// GetUsersList params
// @id: int
// @userData: entity Users
// return entity,error
func (repository *UserRepository) GetUsersList(limit int, offset int) ([]entity.Users, error) {
	users := []entity.Users{}
	query := repository.DB.Table("tb_users")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
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

// CheckEmailUsers params
// @email : string
// @userData: entity Users
// return error
func (repository *UserRepository) CheckEmailUsers(email string, usersData *entity.Users) bool {
	query := repository.DB.Table("tb_users")
	if err := query.Where("email=?", email).First(&usersData).Error; err != nil {
		return false
	}
	return true
}
