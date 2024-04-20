package main

import (
	"github.com/IshaqNiloy/go-rest-api/database"
	"github.com/IshaqNiloy/go-rest-api/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	app := fiber.New()

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
