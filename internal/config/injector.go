// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"learn-go/internal/database/repository"
	"learn-go/internal/module/auth"
	"learn-go/internal/module/packages"
	"learn-go/internal/module/resource"
	"learn-go/internal/module/user"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectDB)

// user wire set
var userSet = wire.NewSet(
	user.UserServiceInit,
	repository.UserRepositoryInit,
	user.UserControllerInit,
	wire.Bind(new(user.UserService), new(*user.UserServiceImpl)),
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Bind(new(user.UserController), new(*user.UserControllerImpl)),
)

// auth wire set
var authSet = wire.NewSet(
	auth.AuthServiceInit,
	auth.AuthControllerInit,
	wire.Bind(new(auth.AuthService), new(*auth.AuthServiceImpl)),
	wire.Bind(new(auth.AuthController), new(*auth.AuthControllerImpl)),
)

// package wire set
var packageSet = wire.NewSet(
	packages.PackageServiceInit,
	repository.PackageRepositoryInit,
	packages.PackageControllerInit,
	wire.Bind(new(packages.PackageService), new(*packages.PackageServiceImpl)),
	wire.Bind(new(repository.PackageRepository), new(*repository.PackageRepositoryImpl)),
	wire.Bind(new(packages.PackageController), new(*packages.PackageControllerImpl)),
)

// resource wire set
var resourceSet = wire.NewSet(
	resource.ResourceServiceInit,
	repository.ResourceRepositoryInit,
	resource.ResourceControllerInit,
	wire.Bind(new(resource.ResourceService), new(*resource.ResourceServiceImpl)),
	wire.Bind(new(repository.ResourceRepository), new(*repository.ResourceRepositoryImpl)),
	wire.Bind(new(resource.ResourceController), new(*resource.ResourceControllerImpl)),
)

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		userSet,
		authSet,
		packageSet,
		resourceSet,
	)
	return nil
}
