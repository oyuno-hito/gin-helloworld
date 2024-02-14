package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
)

type UserController struct{}

func (uc UserController) GET(c *gin.Context) {
	age := float32(1)
	name := "サンプル太郎"
	c.JSON(
		200,
		openapi.UserInfo{
			Age:  &age,
			Name: &name,
		},
	)
}
