package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetUserInfoUseCase(id int) *repository.UserRole {
	// TODO: 秘匿したい情報を環境変数から呼び出すようにする
	// TODO: DIコンテナに登録する
	// TODO: 各種エラーハンドリング
	user := "root"
	password := "root"
	ipAddress := "127.0.0.1"
	port := "3306"
	dbName := "gin-helloworld"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ipAddress, port, dbName)
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	userRepository := repository.UserRepositoryImpl{}
	dto, repositoryErr := userRepository.FindById(db, id)
	fmt.Println(dbErr, repositoryErr)
	// TODO: domainモデルの定義
	return dto
}
