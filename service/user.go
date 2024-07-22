package service

import (
	"github.com/ahay12/api-app/database"
	"github.com/ahay12/api-app/model"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	var users []model.Users
	DB := database.DB
	DB.Find(&users)
	err := ctx.JSON(users)
	if err != nil {
		return err
	}
	return err
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user model.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return ctx.JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(model.Users)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&user)
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user model.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&user)
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user model.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	database.DB.Delete(&user)
	return ctx.JSON(fiber.Map{"message": "User deleted"})
}
