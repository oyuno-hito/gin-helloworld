package main

import (
	"flag"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/database"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"github.com/oyuno-hito/gin-helloworld/src/presentation/controller"
)

type Server struct{}

// GetUserInfo implements openapi.ServerInterface.
func (Server) GetUserInfo(c *gin.Context) {
	userController := controller.UserController{}
	userController.GET(c)
}

// PostLogin implements openapi.ServerInterface.
func (Server) PostLogin(c *gin.Context) {
	postController := controller.LoginController{}
	postController.POST(c)
}

func (Server) PostLogout(c *gin.Context) {
	logoutController := controller.LogoutController{}
	logoutController.POST(c)
}

func main() {
	// port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	r := gin.Default()
	s := Server{}
	options := openapi.GinServerOptions{BaseURL: "api"}

	db := database.NewDb()
	store := gormsessions.NewStore(db, true, []byte("secret"))
	r.Use(sessions.Sessions("session", store))

	openapi.RegisterHandlersWithOptions(r, s, options)

	r.Run(":8080")
}
