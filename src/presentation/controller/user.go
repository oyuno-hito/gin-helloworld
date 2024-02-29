package controller

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
)

type UserController struct {
	userInfoUseCase usecase.UserInfoUseCase
}

func NewUserController(userInfoUseCase usecase.UserInfoUseCase) UserController {
	return UserController{
		userInfoUseCase: userInfoUseCase,
	}
}

func (uc UserController) GET(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	if id == nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New("ログインしてください"))
	}
	model, err := uc.userInfoUseCase.GetUserInfoUseCase(id.(int))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	response := openapi.UserInfo{
		Name: &model.UserName,
		Role: &model.RoleName,
	}
	c.IndentedJSON(http.StatusOK, response)
}
