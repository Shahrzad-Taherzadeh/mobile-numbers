package handler

import (
	"time"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/config" 
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
    tokenString, err := token.SignedString([]byte(config.AppConfig.Server.JWTSecret)) 
}

func JWTMiddleware(c *fiber.Ctx) error {
    token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.AppConfig.Server.JWTSecret), nil 
    })
}