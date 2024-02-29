package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
	"github.com/oyuno-hito/gin-helloworld/src/usecase/model"
)

type LoginController struct {
	loginUseCase usecase.Login
}

func NewLoginController(loginUseCase usecase.Login) LoginController {
	return LoginController{
		loginUseCase: loginUseCase,
	}
}

func (lc LoginController) POST(c *gin.Context) {
	var request openapi.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		return
	}
	session := sessions.Default(c)
	loginUser := model.LoginUser{
		LoginId:  request.LoginId,
		Password: request.Password,
	}
	// TODO: 命名の重複が気になる。login.usecaseとかがワンチャンあるかもしれない
	id := lc.loginUseCase.LoginUseCase(loginUser)
	session.Set("id", id)
	session.Save()
	// TODO セッションIDを返す
	c.JSON(
		200,
		nil,
	)
}
