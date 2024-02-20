package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/presentation/controller"
)

// type Server struct{}

type Server struct {
	UserInfo int64
}

// GetUserInfo implements openapi.ServerInterface.
func (Server) GetUserInfo(c *gin.Context, params openapi.GetUserInfoParams) {
	userController := controller.UserController{}
	userController.GET(c, params)
}

// PostLogin implements openapi.ServerInterface.
func (Server) PostLogin(c *gin.Context) {
	postController := controller.LoginController{}
	postController.POST(c)
}

func NewServerInterface() *Server {
	return &Server{
		UserInfo: 2,
	}
}

func main() {
	// port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	r := gin.Default()
	s := Server{}
	options := openapi.GinServerOptions{BaseURL: "api"}
	openapi.RegisterHandlersWithOptions(r, s, options)

	r.Run(":8080")
}
