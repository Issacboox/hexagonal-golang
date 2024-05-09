package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the Authorization header is present
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Missing token",
			})
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorUnverifiable)
			}
			// Return the secret key used for signing the token
			return []byte("your-secret-key"), nil
		})

		// Check for errors
		if err != nil {
			if err, ok := err.(*jwt.ValidationError); ok {
				if err.Errors&jwt.ValidationErrorMalformed != 0 {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Unauthorized: Malformed token",
					})
				} else if err.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Unauthorized: Token is expired or not valid yet",
					})
				} else {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Unauthorized: Unable to parse token",
					})
				}
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Unable to parse token",
			})
		}

		// Check if the token is valid
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
			})
		}

		// If the token is valid, proceed to the next middleware or route handler
		return c.Next()
	}
}
