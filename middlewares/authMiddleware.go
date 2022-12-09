package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("JWT Cookie")
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Failed Unauthentication",
		})
	}

	return c.Next() // going to the next route
}
