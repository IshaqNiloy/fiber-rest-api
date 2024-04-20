package router

import (
	"github.com/IshaqNiloy/go-rest-api/handler"
	"github.com/IshaqNiloy/go-rest-api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New(), middleware.AuthReq())

	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	//api.Post("/", handler.CreateProduct)
	//api.Delete("/:id", handler.DeleteProduct)
}
