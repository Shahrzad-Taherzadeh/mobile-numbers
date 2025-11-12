package handler

import (
	"time"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}


func JWTMiddleware(c *fiber.Ctx) error {
}