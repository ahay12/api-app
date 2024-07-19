package service

import (
	"api-apps/database"
	"api-apps/model"
	"github.com/gofiber/fiber/v2"
)

func getProducts(ctx *fiber.Ctx) {
	var products []model.Products
	DB := database.DB
	DB.Find(&products)
	err := ctx.JSON(products)
	if err != nil {
		return
	}
}

func getProduct(ctx *fiber.Ctx) {

}

func createProduct(ctx *fiber.Ctx) {

}

func updateProduct(ctx *fiber.Ctx) {

}

func deleteProduct(ctx *fiber.Ctx) {

}
