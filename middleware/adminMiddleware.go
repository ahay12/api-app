package middleware

import (
	"os"
	"strings"

	"github.com/ahay12/api-app/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("API_KEY")

// AdminMiddleware checks if the user has an admin role
func AdminMiddleware(ctx *fiber.Ctx) error {
	// Get the JWT from the Authorization header
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Missing or invalid token", nil, nil)
		return nil
	}

	// Remove "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Unauthorized", nil, []helper.ErrorField{
			{
				ID:      "request",
				Value:   "body",
				Caused:  "token",
				Message: "Invalid token",
			},
		})
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		helper.RespondJSON(ctx, fiber.StatusUnauthorized, "Invalid token", nil, nil)
		return nil
	}

	// Check if the role is admin
	if role, roleExists := claims["role"].(string); !roleExists || role != "admin" {
		helper.RespondJSON(ctx, fiber.StatusForbidden, "Forbidden, only admins can access", nil, nil)
		return nil
	}

	return ctx.Next()
}
