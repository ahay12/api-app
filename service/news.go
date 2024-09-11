package service

import (
	"github.com/ahay12/api-app/database"
	"github.com/ahay12/api-app/helper"
	"github.com/ahay12/api-app/model"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	var Articles []model.News
	DB := database.DB
	DB.Find(&Articles)
	err := ctx.JSON(Articles)
	if err != nil {
		return err
	}
	return err
}

func GetArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Article model.News
	database.DB.First(&Article, id)

	if Article.ID != 0 {
		helper.RespondJSON(ctx, 200, "", Article, nil)
	} else {
		helper.RespondJSON(ctx, 404, "News not found", nil, nil)
	}
	return nil
}

func CreateArticle(ctx *fiber.Ctx) error {
	Article := new(model.News)
	if err := ctx.BodyParser(Article); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&Article)
	return ctx.JSON(Article)
}

func UpdateArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Article model.News
	result := database.DB.First(&Article, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}
	if err := ctx.BodyParser(&Article); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&Article)
	return ctx.JSON(Article)
}

func DeleteArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Article model.News
	result := database.DB.First(&Article, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}
	database.DB.Delete(&Article)
	return ctx.SendStatus(204)
}
