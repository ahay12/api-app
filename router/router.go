package router

import (
	"github.com/ahay12/api-app/service"
	"github.com/gofiber/fiber/v2"
)

func Cors() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return ctx.Next()
	}
}

func Make() *fiber.App {
	app := fiber.New()
	app.Use(Cors())
	v1 := app.Group("/api/v1")
	{
		v1.Get("/products", service.GetProducts)
		v1.Get("/products/:id", service.GetProduct)
		v1.Post("/products", service.CreateProduct)
		v1.Put("/products/:id", service.UpdateProduct)
		v1.Delete("/products/:id", service.DeleteProduct)
	}
	{
		v1.Get("/users", service.GetUsers)
		v1.Get("/users/:id", service.GetUser)
		v1.Post("/users", service.CreateUser)
		v1.Put("/users/:id", service.UpdateUser)
		v1.Delete("/users/:id", service.DeleteUser)
	}

	return app
}
