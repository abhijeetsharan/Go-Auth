package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"github.com/abhijeetsharan/go-auth/database"
	"github.com/abhijeetsharan/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))

	routes.Setup(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
