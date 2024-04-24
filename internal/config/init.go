package config

import (
	"learn-go/internal/database/repository"
	"learn-go/internal/module/auth"
	"learn-go/internal/module/packages"
	"learn-go/internal/module/resource"
	"learn-go/internal/module/user"
)

type Initialization struct {
	// user Initialization
	userRepo repository.UserRepository
	userSvc  user.UserService
	UserCtrl user.UserController

	// auth Initialization
	authSvc  auth.AuthService
	AuthCtrl auth.AuthController

	// package Initialization
	packageRepo repository.PackageRepository
	packageSvc  packages.PackageService
	PackageCtrl packages.PackageController

	// resource Initialization
	ResourceSvc  resource.ResourceService
	ResourceCtrl resource.ResourceController
	ResourceRepo repository.ResourceRepository
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService user.UserService,
	userCtrl user.UserController,
	authSvc auth.AuthService,
	authCtrl auth.AuthController,
	packageRepo repository.PackageRepository,
	packageSvc packages.PackageService,
	PackageCtrl packages.PackageController,

	resourceSvc resource.ResourceService,
	resourceCtrl resource.ResourceController,
	resourceRepo repository.ResourceRepository,

) *Initialization {
	return &Initialization{
		userRepo:    userRepo,
		userSvc:     userService,
		UserCtrl:    userCtrl,
		authSvc:     authSvc,
		AuthCtrl:    authCtrl,
		packageRepo: packageRepo,
		packageSvc:  packageSvc,
		PackageCtrl: PackageCtrl,

		// resource Initialization
		ResourceSvc:  resourceSvc,
		ResourceCtrl: resourceCtrl,
		ResourceRepo: resourceRepo,
	}
}
