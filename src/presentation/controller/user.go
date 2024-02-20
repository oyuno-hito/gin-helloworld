package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
)

type UserController struct{}

func (uc UserController) GET(c *gin.Context, params openapi.GetUserInfoParams) {
	model := usecase.GetUserInfoUseCase(*params.Id)
	response := openapi.UserInfo{
		Name: &model.UserName,
		Role: &model.RoleName,
	}
	c.IndentedJSON(http.StatusOK, response)
}
