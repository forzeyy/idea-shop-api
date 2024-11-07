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
	cartHandler := handlers.NewCartHandler()

	api := app.Group("/api")
	api.Get("/", hello)

	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api.Post("/refresh", authHandler.Refresh)

	api.Get("/products", productHandler.GetAllProducts)
	api.Get("/products/:id", productHandler.GetProductByID)
	api.Get("/products/category/:category_id", productHandler.GetProductsByCategory)
	api.Post("/products", productHandler.CreateProduct)
	api.Post("/upload-url", productHandler.UploadProductImage)

	profile := api.Group("/profile", middleware.Protected())
	profile.Get("/profile", userHandler.GetProfile)
	profile.Patch("/profile", userHandler.UpdateProfile)

	cart := api.Group("/cart", middleware.Protected())
	cart.Get("/", cartHandler.GetCart)
	cart.Post("/add", cartHandler.AddCartItem)
	cart.Patch("item/:id", cartHandler.UpdateCartItem)
	cart.Delete("item/:id", cartHandler.RemoveCartItem)
	cart.Delete("/clear", cartHandler.ClearCart)
}
