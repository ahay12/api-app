package service

import (
	"github.com/ahay12/api-app/database"
	"github.com/ahay12/api-app/helper"
	"github.com/ahay12/api-app/model"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	var articles []model.News
	DB := database.DB
	if err := DB.Find(&articles).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to fetch articles", nil, err.Error())
		return err
	}

	helper.RespondJSON(ctx, fiber.StatusOK, "Articles fetched successfully", articles, nil)
	return nil
}

func GetArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var article model.News
	if err := database.DB.First(&article, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Article not found", nil, err.Error())
	}

	if article.ID != 0 {
		helper.RespondJSON(ctx, fiber.StatusOK, "Success get article", article, nil)
	} else {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "News not found", nil, nil)
	}
	return nil
}

func CreateArticle(ctx *fiber.Ctx) error {
	article := new(model.News)
	if err := ctx.BodyParser(article); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}

	if err := database.DB.Create(&article).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to create article", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully create article", article, nil)
	return nil
}

func UpdateArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var article model.News
	if err := database.DB.First(&article, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusNotFound, "Article not found", nil, err.Error())
		return err
	}

	if err := ctx.BodyParser(&article); err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil, err.Error())
		return err
	}
	if err := database.DB.Save(&article).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to update article", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Successfully update article", article, nil)
	return nil
}

func DeleteArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var article model.News
	if err := database.DB.First(&article, id).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusBadRequest, "Article not found", nil, nil)
		return err
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		helper.RespondJSON(ctx, fiber.StatusInternalServerError, "Failed to delete article", nil, err.Error())
		return err
	}
	helper.RespondJSON(ctx, fiber.StatusOK, "Article deleted successfully", nil, nil)
	return nil
}
