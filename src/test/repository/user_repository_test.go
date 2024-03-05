package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestUserRepository_FindById(t *testing.T) {
	dummyId := 1
	mockDB, mock, _ := NewDbMock()

	t.Run("一致するIDのレコードが存在する場合、紐づくUserRoleが返る", func(t *testing.T) {
		// arrange
		dummyUser := "user"
		dummyRole := "role"
		rows := sqlmock.NewRows([]string{"id", "user_name", "role_name"}).
			AddRow(1, dummyUser, dummyRole)

		query := "SELECT users.name AS user_name, roles.name as role_name FROM `users` left join roles on users.role_id = roles.id WHERE users.id = ?"
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(1).
			WillReturnRows(rows)

		repo := repository.NewUserRepository(mockDB)

		// act
		userRole, err := repo.FindById(dummyId)

		//assert
		expected := &repository.UserRole{
			UserName: dummyUser,
			RoleName: dummyRole,
		}
		assert.Equal(t, expected, userRole)
		assert.Nil(t, err)
	})

	t.Run("一致するIDのレコードが存在しない場合、nilが返る", func(t *testing.T) {
		// arrange
		rows := sqlmock.NewRows([]string{"id", "user_name", "role_name"})

		query := "SELECT users.name AS user_name, roles.name as role_name FROM `users` left join roles on users.role_id = roles.id WHERE users.id = ?"
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(1).
			WillReturnRows(rows)

		repo := repository.NewUserRepository(mockDB)

		// act
		userRole, err := repo.FindById(dummyId)

		//assert
		assert.Nil(t, userRole)
		assert.Nil(t, err)
	})
}
