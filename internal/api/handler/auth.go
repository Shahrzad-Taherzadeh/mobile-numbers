package handler

import (
	"time"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/middleware"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/service"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Name string `json:"name"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	users, _ := service.GetUserList()
	var found *int
	for _, u := range users {
		if u.Name == req.Name {
			id := u.ID
			found = &id
			break
		}
	}

	if found == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
	}

	token, err := middleware.GenerateToken(*found, 24*time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
