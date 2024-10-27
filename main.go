package main

import (
	"log"

	"github.com/forzeyy/idea-shop-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
