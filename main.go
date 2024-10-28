package main

import (
	"log"

	"github.com/forzeyy/idea-shop-api/database"
	"github.com/forzeyy/idea-shop-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
