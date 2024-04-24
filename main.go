package main

import (
	"learn-go/docs"
	"learn-go/internal/config"
	"learn-go/pkg/middleware"
	route "learn-go/route"
	"os"

	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func init() {
	godotenv.Load()
	middleware.InitLog()
	config.ConnectDB()
}

func main() {
	// @title           Gin SAS Service
	// @version         1.0
	// @description     A SAS API in Go using Gin framework.

	// @contact.name   Purjayadi
	// @contact.url    https://www.linkedin.com/in/purjayadi-9a154013a/
	// @contact.email  purjayadi@gmail.com

	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

	// @host      localhost:8080
	// @BasePath  /api
	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	port := os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/api"
	init := config.Init()
	app := route.Init(init)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.PersistAuthorization(true)))

	app.Run(":" + port)
}
