package controllers

import (
	"fmt"

	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/database"
	"github.com/vanessamae23/cvwo/models"
	"github.com/vanessamae23/cvwo/util"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	//Encrypting the password of the user but store as String in DB
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(password)
	fmt.Println(user)

	database.DB.Create(&user)

	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	// checking if the user actually exists
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//Comparing the password user inputs and the password of the user
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// id : strconv.Itoa(int(user.Id))
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id))) //secret key we want to use

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//Cookie
	cookie := fiber.Cookie{
		Name:     "JWT Cookie",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 48),
		HTTPOnly: true, // so that the frontend cannot access this cookie
		SameSite: "None",
		Secure:   true,
	}

	c.Cookie(&cookie)

	// successfully retrieve the user
	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

type Claims struct {
	jwt.StandardClaims // the struct always have the default fields inside standardClaims
}

// Getting the authenticated user
func User(c *fiber.Ctx) error {
	//get the cookie we set
	cookie := c.Cookies("JWT Cookie")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "JWT Cookie",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 48),
		HTTPOnly: true, // setting the expiration time to the past
		
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Cookie revoked successfully",
	})

}
