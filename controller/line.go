package controller

import (
	"strings"

	"github.com/famesensor/playground-go-line-sheet/models"
	"github.com/famesensor/playground-go-line-sheet/ports"
	"github.com/gofiber/fiber/v2"
)

type LineController struct {
	lineSrv ports.LineService
}

func NewLineController(lineSrv ports.LineService) *LineController {
	return &LineController{
		lineSrv,
	}
}

func (ctl *LineController) ReceiveLineMessage(ctx *fiber.Ctx) error {
	body := new(models.LineMessage)

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "body parser error")
	}

	event := body.Events[0]

	if event.Message.Type != "text" {
		go ctl.lineSrv.SendMessageError(event.ReplyToken)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "ok"})
	}

	replyToken := event.ReplyToken
	text := strings.ToLower(event.Message.Text)

	splitText := strings.Split(text, "|")
	if len(splitText) == 3 {
		ctl.lineSrv.AddWord(ctx.Context(), replyToken, splitText[1], splitText[2])
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "ok"})
	}

	switch text {
	case "op":
		go ctl.lineSrv.GetOperations(replyToken)
	case "gw":
		go ctl.lineSrv.GetWords(replyToken)
	case "aw":
		go ctl.lineSrv.AddWordDescription(replyToken)
	default:
		go ctl.lineSrv.SendMessageError(event.ReplyToken)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "ok"})
}
