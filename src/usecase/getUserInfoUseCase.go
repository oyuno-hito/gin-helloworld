package usecase

import (
	"fmt"

	"github.com/oyuno-hito/gin-helloworld/src/service"

	"github.com/oyuno-hito/gin-helloworld/src/repository"
)

type UserInfoUseCase interface {
	GetUserInfoUseCase(id int) (*repository.UserRole, error)
}

type UserInfoUseCaseImpl struct {
	userInfoService service.UserInfoService
}

func NewUserInfoUseCase(userInfoService service.UserInfoService) UserInfoUseCase {
	return UserInfoUseCaseImpl{
		userInfoService: userInfoService,
	}
}

func (uc UserInfoUseCaseImpl) GetUserInfoUseCase(id int) (*repository.UserRole, error) {
	dto, repositoryErr := uc.userInfoService.GetAndValidate(id)
	fmt.Println("repositoryErr: ", repositoryErr)
	// TODO: domainモデルの定義
	return dto, repositoryErr
}
