package routes

import (
	"github.com/forzeyy/idea-shop-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("hello!")
}

func SetupRoutes(app *fiber.App) {
	productHandler := handlers.NewProductHandler()

	api := app.Group("/api")
	api.Get("/", hello)
	//products routes
	api.Get("/products", productHandler.GetAllProducts)
	api.Get("/products/:id", productHandler.GetProductByID)
	api.Post("/products", productHandler.CreateProduct)
	api.Put("/products/:id", productHandler.UpdateProduct)
	api.Delete("/products/:id", productHandler.DeleteProduct)
}
