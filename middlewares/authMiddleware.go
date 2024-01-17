package middlewares

import (
	"github.com/ADEMOLA200/go_admin_application/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized user",
		})
	}

	return c.Next() // Move to the next route
}

