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
		v1.Get("/articles", service.GetArticles)
		v1.Get("/article/:id", service.GetArticle)
		v1.Post("/article", service.CreateArticle)
		v1.Put("/article/:id", service.UpdateArticle)
		v1.Delete("/article/:id", service.DeleteArticle)
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
