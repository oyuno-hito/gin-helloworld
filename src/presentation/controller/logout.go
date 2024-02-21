package controller

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogoutController struct{}

func (lc LogoutController) POST(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println(session)
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	c.JSON(
		200,
		nil,
	)
}
