package apiserver

import (
	"fmt"
	"log"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/handler"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/config" 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	swagger "github.com/swaggo/fiber-swagger" 
)

func Start() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(logger.New())

	app.Get("/swagger/*", swagger.HandlerDefault) 

	app.Post("/login", handler.Login)

	app.Use(handler.JWTMiddleware)

	app.Get("/user", handler.GetUserList)
	app.Get("/user/:id", handler.GetUserByID)
	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUserByID)
	app.Delete("/user/:id", handler.DeleteUserByID)

	app.Post("/user/:id/mobile-number", handler.AddMobileNumber)
	app.Delete("/user/:id/mobile-number/:number", handler.DeleteMobileNumber)

	listenAddr := fmt.Sprintf(":%d", config.AppConfig.Server.Port) 
	log.Printf("Starting server on %s", listenAddr)
	log.Println(app.Listen(listenAddr))
}