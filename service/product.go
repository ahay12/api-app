package service

import (
	"github.com/ahay12/api-app/database"
	"github.com/ahay12/api-app/helper"
	"github.com/ahay12/api-app/model"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(ctx *fiber.Ctx) error {
	var products []model.Products
	DB := database.DB
	DB.Find(&products)
	err := ctx.JSON(products)
	if err != nil {
		return err
	}
	return err
}

func GetProduct(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	var product model.Products
	database.DB.First(&product, id)

	if product.ID != 0 {
		helper.RespondJSON(ctx, 200, "", product, nil)
	} else {
		helper.RespondJSON(ctx, 404, "Product not found", nil, nil)
	}
}

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(model.Products)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&product)
	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product model.Products
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&product)
	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product model.Products
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	database.DB.Delete(&product)
	return ctx.SendStatus(204)
}
