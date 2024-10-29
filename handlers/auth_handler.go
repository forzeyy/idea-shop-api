package handlers

import (
	"github.com/forzeyy/idea-shop-api/auth"
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/forzeyy/idea-shop-api/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
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
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid phone or password",
		})
	}

	if !utils.CheckPassword(user.Password, loginData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid phone or password",
		})
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate token",
		})
	}

	return c.JSON(fiber.Map{"token": token})
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
			"error": "could'nt hash password",
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
