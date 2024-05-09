package handler

import (
	m "bam/internal/app/model"
	s "bam/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthActions interface {
	AuthenticateUser(email, password string) (*m.User, error)
}

type JWTActions interface {
	GenerateToken(userID uint) (string, error)
	VerifyToken(tokenString string) (*s.Claims, error)
}
type AuthHandler struct {
	userService s.UserService
	jwtService  s.JWTActions
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(userService s.UserService, jwtActions s.JWTActions) *AuthHandler {
	return &AuthHandler{userService: userService, jwtService: jwtActions}
}

func (h *AuthHandler) AuthenticateUser(email, password string) (*m.User, error) {
    // Implement your authentication logic here
    user, err := h.userService.AuthenticateUser(email, password)
    if err != nil {
        return nil, err
    }

    // Generate and store the token if needed
    _, _ = h.jwtService.GenerateToken(user.ID)

    return user, nil
}


func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input m.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.userService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logout"})
}
