package routes

import (
	"github.com/ADEMOLA200/go_admin_application/controllers"
	"github.com/ADEMOLA200/go_admin_application/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
}