package main

import (
	"flag"

	"github.com/oyuno-hito/gin-helloworld/src/di"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/oyuno-hito/gin-helloworld/src/database"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
)

func main() {
	// port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	db := database.NewDb()
	r := gin.Default()
	s := di.InitializeServer(db).Server
	options := openapi.GinServerOptions{BaseURL: "api"}

	store := gormsessions.NewStore(db, true, []byte("secret"))
	r.Use(sessions.Sessions("session", store))

	openapi.RegisterHandlersWithOptions(r, s, options)

	r.Run(":8080")
}
