package routes

import (
	"github.com/famesensor/playground-go-line-sheet/controller"
	"github.com/gofiber/fiber/v2"
)

func LineRoutes(app fiber.Router, ctl controller.LineController) {
	r := app.Group("/line")

	r.Post("/webhook", ctl.ReceiveLineMessage)
}
