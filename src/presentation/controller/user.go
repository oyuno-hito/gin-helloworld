package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
)

type UserController struct{}

func (uc UserController) GET(c *gin.Context, params openapi.GetUserInfoParams) {
	session := sessions.Default(c)
	id := session.Get("id")
	model, err := usecase.GetUserInfoUseCase(id.(int))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	response := openapi.UserInfo{
		Name: &model.UserName,
		Role: &model.RoleName,
	}
	c.IndentedJSON(http.StatusOK, response)
}
