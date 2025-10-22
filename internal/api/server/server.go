package apiserver

import (
	"log"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(logger.New())

	app.Post("/login", handler.Login)

	app.Use(handler.JWTMiddleware) 

	app.Get("/user", handler.GetUserList)
	app.Get("/user/:id", handler.GetUserByID)
	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUserByID)
	app.Delete("/user/:id", handler.DeleteUserByID)

	app.Post("/user/:id/mobile-number", handler.AddMobileNumber)
	app.Delete("/user/:id/mobile-number/:number", handler.DeleteMobileNumber)

	log.Println(app.Listen(":8080"))
}
