package routes

import (
	"fiberapi/app/http/handlers/controller"
	"fiberapi/app/http/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	app.Get("/", controller.Index)

	api := app.Group("/api") // /api

	v1 := api.Group("/v1")         // /api/v1
	v1.Get("/users", user.GetUser) // /api/v1/list
	v1.Get("/users/:id", user.GetUserById)
	v1.Get("/verification", user.SendVerification)

}
