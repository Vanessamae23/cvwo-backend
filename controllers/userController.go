package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/database"
	"github.com/vanessamae23/cvwo/models"
	//"github.com/vanessamae23/cvwo/util"
)

// Fetch all users
func AllUsers(c *fiber.Ctx) error {
	var users[]models.User

	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)
	return c.JSON(user)
}