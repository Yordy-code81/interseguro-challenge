package handler

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// GenerateToken godoc
// @Summary Generate an auth token
// @Description Generates a valid JWT token valid for 24h for testing purposes.
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/token [get]
func (h *AuthHandler) GenerateToken(c *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_for_challenge"
	}

	claims := jwt.MapClaims{
		"user": "developer",
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}
