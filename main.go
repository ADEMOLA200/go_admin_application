package main

import (
	"github.com/ADEMOLA200/go_admin_application/database"
	"github.com/ADEMOLA200/go_admin_application/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}