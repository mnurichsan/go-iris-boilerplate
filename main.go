package main

import (
	"iris-boilerplate/config"
	"iris-boilerplate/controller"
	"iris-boilerplate/repository"
	"iris-boilerplate/service"

	"github.com/kataras/iris/v12"
)

func main() {
	//setup configuration
	configuration := config.New()
	database := config.NewDatabase(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	UserController := controller.NewUserController(&userService)
	//setup Iris
	app := iris.New()

	// Example Routing
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON("Hello World")
	})

	//setup routing
	UserController.Route(app)

	app.Listen(":3000")

}
