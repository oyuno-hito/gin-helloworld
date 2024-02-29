package usecase

import (
	"fmt"
	"github.com/oyuno-hito/gin-helloworld/src/service"

	"github.com/oyuno-hito/gin-helloworld/src/repository"
)

type UserInfoUseCase struct {
	userInfoService service.UserInfoService
}

func NewUserInfoUseCase(userInfoService service.UserInfoService) UserInfoUseCase {
	return UserInfoUseCase{
		userInfoService: userInfoService,
	}
}

func (uc UserInfoUseCase) GetUserInfoUseCase(id int) (*repository.UserRole, error) {
	dto, repositoryErr := uc.userInfoService.GetAndValidate(id)
	fmt.Println("repositoryErr: ", repositoryErr)
	// TODO: domainモデルの定義
	return dto, repositoryErr
}
