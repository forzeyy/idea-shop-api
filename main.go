package main

import (
	"log"

	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.ConnectDatabase()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
