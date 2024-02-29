//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/oyuno-hito/gin-helloworld/src/presentation/controller"
	"github.com/oyuno-hito/gin-helloworld/src/repository"
	"github.com/oyuno-hito/gin-helloworld/src/server"
	"github.com/oyuno-hito/gin-helloworld/src/service"
	"github.com/oyuno-hito/gin-helloworld/src/usecase"
	"gorm.io/gorm"
)

var repositorySet = wire.NewSet(
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewUserInfoService,
)

var usecaseSet = wire.NewSet(
	usecase.NewUserInfoUseCase,
	usecase.NewLogin,
)

var controllerSet = wire.NewSet(
	controller.NewLoginController,
	controller.NewLogoutController,
	controller.NewUserController,
)

var serverSet = wire.NewSet(
	server.NewServer,
)

type ServerSet struct {
	Server server.Server
}

func InitializeServer(db *gorm.DB) *ServerSet {
	wire.Build(
		serverSet,
		controllerSet,
		serviceSet,
		usecaseSet,
		repositorySet,
		wire.Struct(new(ServerSet), "*"),
	)
	return nil
}

// type ControllersSet struct {
// 	LoginController  controller.LoginController
// 	LogoutController controller.LogoutController
// 	UserController   controller.UserController
// }

// func InitializeControllers(db *gorm.DB) *ControllersSet {
// 	wire.Build(
// 		repositorySet,
// 		usecaseSet,
// 		controllerSet,
// 		wire.Struct(new(ControllersSet), "*"),
// 	)
// 	return nil
// }
