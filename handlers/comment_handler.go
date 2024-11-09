package handlers

import (
	"github.com/forzeyy/idea-shop-api/models"
	"github.com/forzeyy/idea-shop-api/repositories"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler interface {
	GetCommentsByUser(*fiber.Ctx) error
	GetCommentsByProductID(*fiber.Ctx) error
	NewComment(*fiber.Ctx) error
	DeleteComment(*fiber.Ctx) error
}

type commentHandler struct {
	repo repositories.CommentRepository
}

func NewCommentHandler() CommentHandler {
	return &commentHandler{
		repo: repositories.NewCommentRepository(),
	}
}

func (h *commentHandler) GetCommentsByUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bad request: user not found",
		})
	}

	comment, err := h.repo.GetCommentsByUser(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}

func (h *commentHandler) GetCommentsByProductID(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bad request: product id not found",
		})
	}

	comments, err := h.repo.GetCommentsByProductID(uint(productID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}

func (h *commentHandler) NewComment(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	comment, err = h.repo.CreateComment(comment, uint(productID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}

func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	var comment models.Comment

	commentID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	comment.ID = uint(commentID)
	comment, err = h.repo.DeleteComment(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}
