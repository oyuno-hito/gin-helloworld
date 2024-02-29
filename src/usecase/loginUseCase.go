package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"github.com/oyuno-hito/gin-helloworld/src/usecase/model"
)

type Login struct {
	userRepository repository.UserRepository
}

func NewLogin(userRepository repository.UserRepository) Login {
	return Login{
		userRepository: userRepository,
	}
}

func (login Login) LoginUseCase(loginUser model.LoginUser) *int {
	id, repositoryErr := login.userRepository.FindByLoginInfo(loginUser.LoginId, loginUser.Password)
	fmt.Println("repositoryErr: ", repositoryErr)
	return id
}
