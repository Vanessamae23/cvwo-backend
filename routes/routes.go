package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanessamae23/cvwo/controllers"
	"github.com/vanessamae23/cvwo/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// Need to pass the middleware
	app.Use(middlewares.IsAuthenticated) // needs to be uppercase

	app.Get("/api/user", controllers.User)
	app.Get("/api/user/:id", controllers.GetUser) //Fetching the data of a single user
	app.Get("/api/users", controllers.AllUsers) //Fetching all users
	app.Post("/api/logout", controllers.Logout)

	//Forums functionalities
	app.Get("/api/forum/:id", controllers.GetForum) //Fetching the data of a single user
	app.Get("/api/forums", controllers.AllForums) //Fetching all users
	app.Get("/api/forumsCategory/:category", controllers.AllForumsByCategory) //Fetching all users
	app.Post("/api/forums", controllers.CreateForum)
	app.Put("/api/forums/:id", controllers.UpdateForum)
	app.Delete("/api/forums/:id", controllers.DeleteForum)

	//Comment sections
	app.Post("/api/comment/:forum_id", controllers.CreateComment)
	app.Get("/api/comment/:forum_id", controllers.AllCommentsByForumId)
}
