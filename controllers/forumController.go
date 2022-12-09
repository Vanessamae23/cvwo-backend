package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/database"
	"github.com/vanessamae23/cvwo/models"
	"github.com/vanessamae23/cvwo/util"
)

// Fetch all forums
func AllForums(c *fiber.Ctx) error {
	var forums[]models.Forum

	database.DB.Find(&forums)
	return c.JSON(forums)
}

func AllForumsByCategory(c *fiber.Ctx) error {
	var forums[]models.Forum
	category := c.Params("category")

	database.DB.Where("category = ?", category).Find(&forums)
	
	return c.JSON(forums)
}

func CreateForum(c *fiber.Ctx) error {
	cookie := c.Cookies("JWT Cookie")
	id, _ := util.ParseJwt(cookie)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	var forum models.Forum

	if err := c.BodyParser(&forum); err != nil {
		return err
	}
	forum.UserId = user.Id
	forum.Username = user.Name
	database.DB.Create(&forum)
	return c.JSON(forum)
}

func GetForum(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	forum := models.Forum{
		Id: uint(id),
	}

	database.DB.Find(&forum)
	return c.JSON(forum)
}

func UpdateForum(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	forum := models.Forum{
		Id: uint(id),
	}

	if err := c.BodyParser(&forum); err != nil {
		return err
	}

	database.DB.Model(&forum).Updates(forum)
	return c.JSON(forum)
}

func DeleteForum(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	forum := models.Forum{
		Id: uint(id),
	}
	var comments models.Comment

	database.DB.Where("forum_id", id).Delete(&comments)
	database.DB.Delete(&forum)
	
	return nil
}

