package router

import "github.com/gofiber/fiber/v2"

func Make() *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/products", getProducts, func(ctx *fiber.Ctx) error {
		panic("this panic if caught by fiber")
	})
	app.Get("/api/v1/products/:id", getProduct)
	app.Post("/api/v1/products", createProduct)
	app.Put("/api/v1/products/:id", updateProduct)
	app.Delete("/api/v1/products/:id", deleteProduct)
	return app
}
