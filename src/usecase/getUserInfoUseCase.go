package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/repository"
)

type UserInfoUseCase struct {
	userRepository repository.UserRepository
}

func NewUserInfoUseCase(userRepository repository.UserRepository) UserInfoUseCase {
	return UserInfoUseCase{
		userRepository: userRepository,
	}
}

func (uc UserInfoUseCase) GetUserInfoUseCase(id int) (*repository.UserRole, error) {
	dto, repositoryErr := uc.userRepository.FindById(id)
	fmt.Println("repositoryErr: ", repositoryErr)
	// TODO: domainモデルの定義
	return dto, repositoryErr
}
