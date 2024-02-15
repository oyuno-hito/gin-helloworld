package controller

import "github.com/gin-gonic/gin"

type LoginController struct{}

func (lc LoginController) POST(c *gin.Context) {
	// p:=
	c.JSON(
		200,
		nil,
	)
}
