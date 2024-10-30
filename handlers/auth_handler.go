package handlers

import (
	"github.com/forzeyy/idea-shop-api/auth"
	"github.com/forzeyy/idea-shop-api/middleware"
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/forzeyy/idea-shop-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
	RefreshToken(*fiber.Ctx) error
}

type authHandler struct {
	userRepo repositories.UserRepository
}

func NewAuthHandler() AuthHandler {
	return &authHandler{
		userRepo: repositories.NewUserRepository(),
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var loginData struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	user, err := h.userRepo.GetUserByPhone(loginData.Phone)
	if err != nil || !utils.CheckPassword(user.Password, loginData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid phone or password",
		})
	}

	accessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate token",
		})
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate token",
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't hash password",
		})
	}
	user.Password = hashedPassword

	if _, err := h.userRepo.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *authHandler) RefreshToken(c *fiber.Ctx) error {
	var inputToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&inputToken); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bad input",
		})
	}

	token, err := jwt.Parse(inputToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return middleware.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "expired or invalid refresh token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	userID := uint(claims["user_id"].(float64))
	accessToken, err := auth.GenerateAccessToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate access token",
		})
	}

	return c.JSON(fiber.Map{
		"access_token": accessToken,
	})
}
