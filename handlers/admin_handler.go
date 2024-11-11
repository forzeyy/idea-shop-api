package handlers

import (
	"os"
	"strconv"

	"github.com/forzeyy/idea-shop-api/auth"
	"github.com/forzeyy/idea-shop-api/middleware"
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/forzeyy/idea-shop-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AdminHandler interface {
	UploadProductImage(*fiber.Ctx) error
	AdminLogin(*fiber.Ctx) error
	AdminRefresh(*fiber.Ctx) error
	AdminRegister(*fiber.Ctx) error
	AddCategory(*fiber.Ctx) error
}

type adminHandler struct {
	adminRepo    repositories.AdminRepository
	S3Service    *middleware.S3Service
	productRepo  repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

func NewAdminHandler() AdminHandler {
	return &adminHandler{
		adminRepo:    repositories.NewAdminRepository(),
		S3Service:    middleware.NewS3Service(utils.S3Client, os.Getenv("BUCKET_NAME")),
		productRepo:  repositories.NewProductRepository(),
		categoryRepo: repositories.NewCategoryRepository(),
	}
}

func (h *adminHandler) UploadProductImage(c *fiber.Ctx) error {
	stringProductID := c.FormValue("product_id")

	productID, err := strconv.Atoi(stringProductID)
	if err != nil || stringProductID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "uncorrect product_id value",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open uploaded file",
		})
	}

	defer fileContent.Close()

	URL, err := h.S3Service.UploadImageToS3(c.Context(), uint(productID), fileContent, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.productRepo.UpdateProductImageURL(uint(productID), URL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"product_id": productID,
		"upload_url": URL,
	})
}

func (h *adminHandler) AdminLogin(c *fiber.Ctx) error {
	var loginData struct {
		AdminName string `json:"admin_name"`
		Password  string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	admin, err := h.adminRepo.GetAdminByName(loginData.AdminName)
	if err != nil || !utils.CheckPassword(admin.Password, loginData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid phone or password",
		})
	}

	accessToken, err := auth.GenerateAccessToken(admin.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate token",
		})
	}

	refreshToken, err := auth.GenerateRefreshToken(admin.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "admin_access_token",
		Value:    accessToken,
		Expires:  auth.AccessTokenLifetime,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "admin_refresh_token",
		Value:    refreshToken,
		Expires:  auth.RefreshTokenLifetime,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func (h *adminHandler) AdminRefresh(c *fiber.Ctx) error {
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

	adminID := uint(claims["admin_id"].(float64))
	accessToken, err := auth.GenerateAccessToken(adminID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't generate access token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "admin_access_token",
		Value:    accessToken,
		Expires:  auth.AccessTokenLifetime,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "access token successfully refreshed",
	})
}

func (h *adminHandler) AdminRegister(c *fiber.Ctx) error {
	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't hash password",
		})
	}
	admin.Password = hashedPassword

	if _, err := h.adminRepo.CreateAdmin(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't create admin",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *adminHandler) AddCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := h.categoryRepo.CreateCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}
