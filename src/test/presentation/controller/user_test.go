package controller

import (
	"encoding/json"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mocks "github.com/oyuno-hito/gin-helloworld/mocks/usecase"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/presentation/controller"
)

func TestUserController_GET(t *testing.T) {
	dummyId := 1
	userControllerGetPath := "/api/user_info"
	t.Run("正常系の場合ステータス200とUserInfoモデルが返ってくること",
		func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserUsecase := mocks.NewMockUserInfoUseCase(ctrl)
			dummyUserName := "dummyName"
			dummyRole := "dummyRole"
			mockUserUsecase.EXPECT().GetUserInfoUseCase(dummyId).Return(
				&repository.UserRole{
					UserName: dummyUserName,
					RoleName: dummyRole,
				},
				nil,
			)

			requestBody := strings.NewReader("")

			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)

			c.Request, _ = http.NewRequest(
				http.MethodGet,
				userControllerGetPath,
				requestBody,
			)
			c.Set("session", dummyId)

			// act
			userController := controller.NewUserController(mockUserUsecase)
			userController.GET(c)

			// assert
			var expected map[string]interface{}
			expectedBytes, _ := json.Marshal(openapi.UserInfo{
				Name: &dummyUserName,
				Role: &dummyRole,
			})
			_ = json.Unmarshal(expectedBytes, &expected)
			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Regexp(t, 200, response.Code)
			assert.Equal(t, expected, responseBody)
		})

	t.Run("セッション情報がcontextに保持されていない場合BadRequestが返ること", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserUsecase := mocks.NewMockUserInfoUseCase(ctrl)

		requestBody := strings.NewReader("")

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)

		c.Request, _ = http.NewRequest(
			http.MethodGet,
			userControllerGetPath,
			requestBody,
		)

		// act
		userController := controller.NewUserController(mockUserUsecase)
		userController.GET(c)

		// assert
		assert.Equal(t, 400, response.Code)
	})
}
