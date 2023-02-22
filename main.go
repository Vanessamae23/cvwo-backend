package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vanessamae23/cvwo/database"
	"github.com/vanessamae23/cvwo/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //frontend to get the cookie
	})) // cors will block req from diff ports

	routes.Setup(app)

	app.Listen(":8000")

}
