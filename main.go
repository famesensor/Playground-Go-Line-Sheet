package main

import (
	"context"
	"log"

	"github.com/famesensor/playground-go-line-sheet/configs"
	"github.com/famesensor/playground-go-line-sheet/controller"
	"github.com/famesensor/playground-go-line-sheet/pkg/botline"
	"github.com/famesensor/playground-go-line-sheet/pkg/sheet"
	"github.com/famesensor/playground-go-line-sheet/repositories"
	"github.com/famesensor/playground-go-line-sheet/routes"
	"github.com/famesensor/playground-go-line-sheet/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	configs.InitViper()
	config := configs.GetConfig()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())

	// TODO: handle error
	lintBot, _ := botline.ConnectLineBot(config)
	sheetConn, _ := sheet.ConnectGoogleSheet(ctx, config)

	// repositories
	lineRepo := repositories.NewLineRepository(lintBot)
	sheetRepo := repositories.NewSheetRepository(sheetConn)

	// service
	lineSrv := services.NewLineService(lineRepo, sheetRepo, config.SheetSpreadSheetId)

	// controller
	lineCtl := controller.NewLineController(lineSrv)

	// routes
	routes.HealthRoutes(app)
	routes.LineRoutes(app, *lineCtl)

	if err := app.Listen(":" + config.AppPort); err != nil {
		log.Fatal(err)
	}
}
