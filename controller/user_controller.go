package controller

import (
	"iris-boilerplate/service"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(app *iris.Application) {
	app.Get("api/user", controller.List)
}

func (controller *UserController) List(ctx iris.Context) {
	responses := controller.UserService.List()

	ctx.JSON(responses)
}
