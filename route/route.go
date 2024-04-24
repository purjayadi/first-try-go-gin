package router

import (
	"learn-go/internal/config"
	"learn-go/pkg/dto"
	"learn-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		// public route
		// auth route
		auth := api.Group("/auth")
		auth.POST("/login", middleware.ValidateInputMiddleware(&dto.LoginDto{}), init.AuthCtrl.Login)
		auth.POST("/register", middleware.ValidateInputMiddleware(&dto.CreateUserDto{}), init.AuthCtrl.Register)

		// private route
		api.Use(middleware.AuthMiddleware())
		// user route
		user := api.Group("/user")
		user.GET("", middleware.ValidateInputMiddleware(&dto.GetUserDto{}), init.UserCtrl.GetAllUser)
		user.GET("/:email", init.UserCtrl.GetUserByEmail)
		user.POST("", middleware.ValidateInputMiddleware(&dto.CreateUserDto{}), init.UserCtrl.AddNewUser)
		user.PATCH("/:id", middleware.ValidateInputMiddleware(&dto.UpdateUserDto{}), init.UserCtrl.UpdateUser)
		user.DELETE("/:id", init.UserCtrl.DeleteUser)

		// package route
		packages := api.Group("/package")
		packages.GET("", middleware.ValidateInputMiddleware(&dto.GetPackageDto{}), init.PackageCtrl.GetAllPackage)
		packages.GET("/detail/:id", init.PackageCtrl.GetPackageById)
		packages.POST("", middleware.ValidateInputMiddleware(&dto.PackageDto{}), init.PackageCtrl.AddNewPackage)
		packages.PATCH("/:id", middleware.ValidateInputMiddleware(&dto.UpdatePackageDto{}), init.PackageCtrl.UpdatePackage)
		packages.DELETE("/:id", init.PackageCtrl.DeletePackage)

		// resource route
		resources := api.Group("/resource")
		resources.GET("", middleware.ValidateInputMiddleware(&dto.SearchResourceDto{}), init.ResourceCtrl.GetAllResource)
		resources.GET("/detail/:id", init.ResourceCtrl.GetResourceById)
		resources.POST("", middleware.ValidateInputMiddleware(&dto.ResourceDto{}), init.ResourceCtrl.AddNewResource)
		resources.PATCH("/:id", middleware.ValidateInputMiddleware(&dto.UpdateResourceDto{}), init.ResourceCtrl.UpdateResource)
		resources.DELETE("/:id", init.ResourceCtrl.DeleteResource)

	}

	return router
}
