package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	mocks "github.com/oyuno-hito/gin-helloworld/mocks/repository"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserInfoService_GetAndValidate(t *testing.T) {
	dummyUserName := "dummyName"
	dummyRole := "dummyRole"
	dummyId := 1

	dummyUserRole := repository.UserRole{
		UserName: dummyUserName,
		RoleName: dummyRole,
	}
	t.Run(
		"repositoryからUserRoleが返ってくる場合その値がそのまま返ること",
		func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			mockUserRepository.EXPECT().FindById(dummyId).Return(
				&dummyUserRole, nil,
			)

			userInfoService := UserInfoServiceImpl{
				userRepository: mockUserRepository,
			}

			expectedValue := &repository.UserRole{
				UserName: dummyUserName,
				RoleName: dummyRole,
			}

			// act
			actualValue, actualErr := userInfoService.GetAndValidate(dummyId)

			// assert
			assert.Equal(t, expectedValue, actualValue)
			assert.Nil(t, actualErr)
		},
	)
	t.Run(
		"repositoryからerrorが返ってくる場合UserRoleがnilを返すこと",
		func(t *testing.T) {
			// arrange
			dummyError := errors.New("dummy")

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepository := mocks.NewMockUserRepository(ctrl)
			// NOTE: モックされているがrepositoryの実装としては排他的に値を返す実装になっているためこのモックは微妙かもしれない
			mockUserRepository.EXPECT().FindById(dummyId).Return(
				&dummyUserRole, dummyError,
			)

			userInfoService := UserInfoServiceImpl{
				userRepository: mockUserRepository,
			}

			// act
			actualValue, actualErr := userInfoService.GetAndValidate(dummyId)

			// assert
			assert.Nil(t, actualValue)
			assert.Equal(t, dummyError, actualErr)
		},
	)

}
