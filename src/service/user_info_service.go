package service

import "github.com/oyuno-hito/gin-helloworld/src/repository"

type UserInfoService interface {
	GetAndValidate(id int) (*repository.UserRole, error)
}

type UserInfoServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserInfoService(userRepository repository.UserRepository) UserInfoService {
	return &UserInfoServiceImpl{
		userRepository: userRepository,
	}
}

func (userInfoService *UserInfoServiceImpl) GetAndValidate(id int) (*repository.UserRole, error) {
	userRole, repositoryErr := userInfoService.userRepository.FindById(id)
	// TODO: 現状repositoryErrが返ってくるわけではなくormのエラーが吐かれてそこで処理が落ちているため修正する
	if repositoryErr != nil {
		return nil, repositoryErr
	}
	return userRole, nil
}
