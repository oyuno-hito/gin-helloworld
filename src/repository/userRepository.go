package repository

import (
	"errors"

	"github.com/oyuno-hito/gin-helloworld/src/repository/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(db *gorm.DB, id int) (*dto.User, error)
	FindByLoginIdOrNull(db *gorm.DB, id int) (*int, error)
	FindByLoginId(db *gorm.DB, id int) (int, error)
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

func (u *UserRepositoryImpl) FindByLoginInfoOrNull(db *gorm.DB, loginId string, password string) (*int, error) {
	user := dto.User{
		Login_id: loginId,
		Password: password,
	}
	result := db.Where(&user).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.Id, nil
}

func (u *UserRepositoryImpl) FindByLoginInfo(db *gorm.DB, loginId string, password string) (*int, error) {
	id, err := u.FindByLoginInfoOrNull(db, loginId, password)

	// TODO: エラーログ設計
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, errors.New("ログインに失敗しました")
	}
	return id, nil
}
