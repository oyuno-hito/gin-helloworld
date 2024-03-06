package main

import (
	"flag"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/gin-middleware"
	"github.com/oyuno-hito/gin-helloworld/src/database"
	"github.com/oyuno-hito/gin-helloworld/src/di"
	"github.com/oyuno-hito/gin-helloworld/src/middleware"
	"github.com/oyuno-hito/gin-helloworld/src/openapi"
	"log"
)

func main() {
	// port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	db := database.NewDb()
	r := gin.Default()
	s := di.InitializeServer(db).Server
	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger spec: %v\n", err)
	}

	r.Use(ginmiddleware.OapiRequestValidator(swagger))

	serverOptions := openapi.GinServerOptions{
		BaseURL: "api",
		Middlewares: []openapi.MiddlewareFunc{
			openapi.MiddlewareFunc(middleware.LoginCheckMiddleware(db, swagger)),
		},
	}

	store := gormsessions.NewStore(db, true, []byte("secret"))
	r.Use(sessions.Sessions("session", store))

	openapi.RegisterHandlersWithOptions(r, s, serverOptions)

	r.Run(":8080")
}
