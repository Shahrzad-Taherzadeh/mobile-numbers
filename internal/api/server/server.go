package apiserver

import (
	"log"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()
	SetGlobalHandlers(app)
	SetupRoutes(app)
	log.Println(app.Listen(":8080"))
}

func SetGlobalHandlers(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(logger.New())
	app.Use(handler.ValidateJwtToken)

}
