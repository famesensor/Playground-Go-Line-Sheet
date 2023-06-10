package routes

import "github.com/gofiber/fiber/v2"

func HealthRoutes(app fiber.Router) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "ok"})
	})
}
