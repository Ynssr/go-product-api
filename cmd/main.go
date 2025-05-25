package main

import (
	"Product/config"
	"Product/handler"
	"Product/repository"
	"Product/router"
	"Product/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	repo := repository.NewProductRepository(config.ProductCollection)
	svc := service.NewProductService(repo)
	handler := handler.NewProductHandler(svc)

	router.SetupRoutes(app, handler)

	app.Listen(":3000")
}
