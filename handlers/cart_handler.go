package handlers

import (
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/gofiber/fiber/v2"
)

type CartHandler interface {
	GetCart(*fiber.Ctx) error
	AddCartItem(*fiber.Ctx) error
	UpdateCartItem(*fiber.Ctx) error
	RemoveCartItem(*fiber.Ctx) error
	ClearCart(*fiber.Ctx) error
}

type cartHandler struct {
	repo repositories.CartRepository
}

func NewCartHandler() CartHandler {
	return &cartHandler{
		repo: repositories.NewCartRepository(),
	}
}

func (h *cartHandler) GetCart(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	cart, err := h.repo.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}

func (h *cartHandler) AddCartItem(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	cart, err := h.repo.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	item, err = h.repo.AddItemToCart(cart.ID, item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *cartHandler) UpdateCartItem(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	cart, err := h.repo.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	itemID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	item.ID = uint(itemID)
	item, err = h.repo.UpdateCartItem(cart.ID, item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *cartHandler) RemoveCartItem(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	cart, err := h.repo.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var item models.CartItem
	itemID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	item.ID = uint(itemID)
	item, err = h.repo.RemoveCartItem(cart.ID, item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *cartHandler) ClearCart(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	cart, err := h.repo.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = h.repo.ClearCart(cart.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cart cleared",
	})
}
