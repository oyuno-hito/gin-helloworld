package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
)

type LoginController struct{}

func (lc LoginController) POST(c *gin.Context) {
	var request openapi.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		return
	}
	session := sessions.Default(c)
	id := usecase.LoginUseCase(request.LoginId, request.Password)
	session.Set("id", id)
	session.Save()
	// TODO セッションIDを返す
	c.JSON(
		200,
		nil,
	)
}
