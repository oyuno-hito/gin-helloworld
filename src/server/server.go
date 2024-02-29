package server

import (
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/presentation/controller"
)

type Server struct {
	userController   controller.UserController
	loginController  controller.LoginController
	logoutController controller.LogoutController
}

func NewServer(
	userController controller.UserController,
	loginController controller.LoginController,
	logoutController controller.LogoutController,
) Server {
	return Server{
		userController:   userController,
		loginController:  loginController,
		logoutController: logoutController,
	}
}

// GetUserInfo implements openapi.ServerInterface.
func (server Server) GetUserInfo(c *gin.Context) {
	server.userController.GET(c)
}

// PostLogin implements openapi.ServerInterface.
func (server Server) PostLogin(c *gin.Context) {
	server.loginController.POST(c)
}

func (server Server) PostLogout(c *gin.Context) {
	server.logoutController.POST(c)
}
