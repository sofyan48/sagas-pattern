package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/svc_user/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_user/src/utils/database"
)

// LoginRepository types
type LoginRepository struct {
	DB gorm.DB
}

// LoginRepositoryHandler Login handler repo
// return: LoginRepository
func LoginRepositoryHandler() *LoginRepository {
	return &LoginRepository{DB: *database.GetTransactionConnection()}
}

// LoginRepositoryInterface interface
type LoginRepositoryInterface interface {
	GetLoginByID(id int, loginData *entity.Login) error
	GetLoginList(limit int, offset int) ([]entity.Login, error)
	GetLoginByUsername(username string, loginData *entity.Login) error
}

// GetLoginByID params
// @id: int
// @loginData: entity Login
// wg *sync.WaitGroup
// return error
func (repository *LoginRepository) GetLoginByID(id int, loginData *entity.Login) error {
	query := repository.DB.Table("tb_login")
	query = query.Where("id=? AND deleted_at IS NULL", id)
	query = query.First(&loginData)
	return query.Error
}

// GetLoginByUsername params
// @id: int
// @loginData: entity Login
// wg *sync.WaitGroup
// return error
func (repository *LoginRepository) GetLoginByUsername(username string, loginData *entity.Login) error {
	query := repository.DB.Table("tb_login")
	query = query.Where("username=? AND deleted_at IS NULL", username)
	query = query.First(&loginData)
	return query.Error
}

// GetLoginList params
// @id: int
// @loginData: entity Login
// return entity,error
func (repository *LoginRepository) GetLoginList(limit int, offset int) ([]entity.Login, error) {
	users := []entity.Login{}
	query := repository.DB.Table("tb_login")
	query = query.Limit(limit).Offset(offset)
	query = query.Where("deleted_at IS NULL")
	query = query.Find(&users)
	return users, query.Error
}
