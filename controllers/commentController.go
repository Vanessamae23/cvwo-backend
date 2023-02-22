package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/database"
	"github.com/vanessamae23/cvwo/models"
	"github.com/vanessamae23/cvwo/util"
)

func CreateComment(c *fiber.Ctx) error {
	cookie := c.Cookies("JWT Cookie")
	id, _ := util.ParseJwt(cookie)
	var user models.User
	var forum models.Forum
	forum_id, _ := strconv.Atoi(c.Params("forum_id"))
	database.DB.Where("id = ?", id).First(&user)
	database.DB.Where("id = ?", forum_id).First(&forum)
	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		return err
	}
	comment.UserId = user.Id
	comment.Username = user.Name
	comment.ForumId = forum.Id
	database.DB.Create(&comment)
	return c.JSON(comment)
}


func AllCommentsByForumId(c *fiber.Ctx) error {
	var comments[]models.Comment
	forum_id, _ := strconv.Atoi(c.Params("forum_id"))

	database.DB.Where("forum_id = ?", forum_id).Find(&comments)
	return c.JSON(comments)
}

func UpdateComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	comment := models.Comment{
		Id: uint(id),
	}

	if err := c.BodyParser(&comment); err != nil {
		return err
	}

	database.DB.Model(&comment).Updates(comment)
	return c.JSON(comment)
}

func DeleteComment(c *fiber.Ctx) error {
	fmt.Println("SDSSDDS")
	id, _ := strconv.Atoi(c.Params("id"))
	comment := models.Comment{
		Id: uint(id),
	}
	print(&comment)	
	database.DB.Delete(&comment)
	return nil
}

