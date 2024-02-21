package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/database"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
)

func GetUserInfoUseCase(id int) (*repository.UserRole, error) {
	db := database.NewDb()
	userRepository := repository.UserRepositoryImpl{}
	dto, repositoryErr := userRepository.FindById(db, id)
	fmt.Println("repositoryErr: ", repositoryErr)
	// TODO: domainモデルの定義
	return dto, repositoryErr
}
