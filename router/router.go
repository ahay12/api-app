package router

import (
	"github.com/ahay12/api-app/middleware"
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

	// Public routes
	{
		v1.Get("/articles", service.GetArticles)
		v1.Get("/article/:id", service.GetArticle)
	}

	{

		v1.Post("/article", middleware.AdminMiddleware, service.CreateArticle)
		v1.Put("/article/:id", middleware.AdminMiddleware, service.UpdateArticle)
		v1.Delete("/article/:id", middleware.AdminMiddleware, service.DeleteArticle)
	}
	{
		v1.Get("/users", middleware.AdminMiddleware, service.GetUsers)
		v1.Get("/user/:id", middleware.AdminMiddleware, service.GetUser)
		v1.Post("/user", service.CreateUser)
		v1.Put("/user/:id", middleware.AdminMiddleware, service.UpdateUser)
		v1.Delete("/user/:id", middleware.AdminMiddleware, service.DeleteUser)
	}
	{
		v1.Post("/login", service.Login)
	}

	return app
}
