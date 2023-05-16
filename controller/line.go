package controller

import "github.com/gofiber/fiber/v2"

type LineController struct {
}

func NewLineController() *LineController {
	return &LineController{}
}

func (ctl *LineController) ReceiveLineMessage(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "ok"})
}
