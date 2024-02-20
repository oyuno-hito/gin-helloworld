package repository

import (
	"github.com/oyuno-hito/gin-helloworld/src/repository/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(db *gorm.DB, id int) (*dto.User, error)
}

type UserRepositoryImpl struct{}

type UserRole struct {
	UserName string
	RoleName string
}

func (u *UserRepositoryImpl) FindById(db *gorm.DB, id int) (*UserRole, error) {
	var userRole UserRole
	fields := "users.name AS user_name, roles.name as role_name"
	query := db.Table("users").Select(fields).Where("users.id = ?", id).Joins("left join roles on users.role_id = roles.id").Scan(&userRole)
	if err := query.Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}
