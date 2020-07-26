package main

import (
	"boilerplate/database"
	"boilerplate/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

var (
	port = flag.Int("port", 3001, "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New()
	app.Settings.Prefork = *prod // go run app.go -prod

	// Middleware
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=3000
}
