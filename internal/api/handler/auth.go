package handler

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if input.Username != "admin" || input.Password != "1234" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  input.Username,
		"exp":       time.Now().Add(time.Hour * 2).Unix(),
		"client_ip": c.IP(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token creation failed"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func ValidateJwtToken(c *fiber.Ctx) error {
	var (
		token *jwt.Token
		err   error
	)

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	if len(jwtToken) != 0 {
		token, err = jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed token"})
	}

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	return c.Next()
}
