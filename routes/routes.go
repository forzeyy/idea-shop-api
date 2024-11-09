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
	adminHandler := handlers.NewAdminHandler()
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()
	userHandler := handlers.NewUserHandler()
	cartHandler := handlers.NewCartHandler()
	commentHandler := handlers.NewCommentHandler()

	app.Post("/admin/login", adminHandler.AdminLogin)

	admin := app.Group("/admin", middleware.Protected())
	admin.Post("/register", adminHandler.AdminRegister)
	admin.Post("/refresh", adminHandler.AdminRefresh)
	admin.Post("/uploadimage", adminHandler.UploadProductImage)
	admin.Post("/products", productHandler.CreateProduct)

	api := app.Group("/api")
	api.Get("/", hello)

	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api.Post("/refresh", authHandler.Refresh)

	products := api.Group("/products")
	products.Get("/", productHandler.GetAllProducts)
	products.Get("/:id", productHandler.GetProductByID)
	products.Get("/category/:category_id", productHandler.GetProductsByCategoryID)
	products.Get("/comments", commentHandler.GetCommentsByProductID)

	profile := api.Group("/profile", middleware.Protected())
	profile.Get("/profile", userHandler.GetProfile)
	profile.Patch("/profile", userHandler.UpdateProfile)

	cart := api.Group("/cart", middleware.Protected())
	cart.Get("/", cartHandler.GetCart)
	cart.Post("/add", cartHandler.AddCartItem)
	cart.Patch("item/:id", cartHandler.UpdateCartItem)
	cart.Delete("item/:id", cartHandler.RemoveCartItem)
	cart.Delete("/clear", cartHandler.ClearCart)

	comment := api.Group("/comment", middleware.Protected())
	comment.Get("/:user_id", commentHandler.GetCommentsByUser)
	comment.Post("/:product_id", commentHandler.NewComment)
	comment.Delete("/:id", commentHandler.DeleteComment)
}
