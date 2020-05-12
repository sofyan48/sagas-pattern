package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_user/src/app/v1/api/user/entity"
	"github.com/sofyan48/svc_user/src/utils/database"
)

// UserRepository types
type UserRepository struct {
	DB gorm.DB
}

// UserRepositoryHandler User handler repo
// return: UserRepository
func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *database.GetTransactionConnection()}
}

// UserRepositoryInterface interface
type UserRepositoryInterface interface {
	GetUserByID(id string, loginData *entity.Users) error
	GetUserList(limit int, offset int) ([]entity.Users, error)
}

// GetUserByID params
// @id: int
// @loginData: entity User
// wg *sync.WaitGroup
// return error
func (repository *UserRepository) GetUserByID(id string, loginData *entity.Users) error {
	query := repository.DB.Table("tb_users")
	query = query.Where("id_user=? AND deleted_at IS NULL", id)
	query = query.First(&loginData)
	return query.Error
}

// GetUserList params
// @id: int
// @loginData: entity User
// return entity,error
func (repository *UserRepository) GetUserList(limit int, offset int) ([]entity.Users, error) {
	users := []entity.Users{}
	query := repository.DB.Table("tb_users")
	query = query.Limit(limit).Offset(offset)
	query = query.Where("deleted_at IS NULL")
	query = query.Find(&users)
	return users, query.Error
}
