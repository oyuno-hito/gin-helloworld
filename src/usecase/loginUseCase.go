package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/database"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
)

func LoginUseCase(loginId string, password string) *int {
	db := database.NewDb()
	userRepository := repository.UserRepositoryImpl{}
	id, repositoryErr := userRepository.FindByLoginInfo(db, loginId, password)
	fmt.Println("repositoryErr: ", repositoryErr)
	return id
}
