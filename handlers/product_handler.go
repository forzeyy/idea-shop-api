package handlers

import (
	"os"

	"github.com/forzeyy/idea-shop-api/middleware"
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/forzeyy/idea-shop-api/utils"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	UploadProductImage(*fiber.Ctx) error
	GetAllProducts(*fiber.Ctx) error
	GetProductByID(*fiber.Ctx) error
	GetProductsByCategory(*fiber.Ctx) error
	CreateProduct(*fiber.Ctx) error
	UpdateProduct(*fiber.Ctx) error
	DeleteProduct(*fiber.Ctx) error
}

type productHandler struct {
	repo      repositories.ProductRepository
	S3Service *middleware.S3Service
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo:      repositories.NewProductRepository(),
		S3Service: middleware.NewS3Service(utils.S3Client, os.Getenv("BUCKET_NAME")),
	}
}

func (h *productHandler) UploadProductImage(c *fiber.Ctx) error {
	var req struct {
		ProductID uint   `json:"product_id"`
		Filename  string `json:"filename"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	uploadURL, err := h.S3Service.GenerateUploadURL(req.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate upload URL",
		})
	}

	if err := h.repo.UpdateProductImageURL(req.ProductID, req.Filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "couldn't update product image url",
		})
	}

	return c.JSON(fiber.Map{"upload_url": uploadURL})
}

func (h *productHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (h *productHandler) GetProductByID(c *fiber.Ctx) error {
	prodID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err := h.repo.GetProductByID(uint(prodID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *productHandler) GetProductsByCategory(c *fiber.Ctx) error {
	categoryID, err := c.ParamsInt("category_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	products, err := h.repo.GetProductsByCategory(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err := h.repo.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	prodID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.ID = uint(prodID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	prodID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.ID = uint(prodID)
	product, err = h.repo.DeleteProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
