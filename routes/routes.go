package routes

import (
	"github.com/forzeyy/idea-shop-api/handlers"
	"github.com/forzeyy/idea-shop-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("hello!")
}

func SetupRoutes(app *fiber.App) {
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()
	userHandler := handlers.NewUserHandler()

	api := app.Group("/api")
	api.Get("/", hello)

	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api.Post("/refresh", authHandler.Refresh)

	api.Get("/products", productHandler.GetAllProducts)
	api.Get("/products/:id", productHandler.GetProductByID)

	protected := api.Group("/protected", middleware.Protected())

	// more protected routes later
	protected.Get("/profile", userHandler.GetProfile)
	protected.Put("/profile", userHandler.UpdateProfile)
}
