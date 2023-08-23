package main

import (
	"learnfiber/controllers"
	"learnfiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()

	controller := controllers.NewHandler()

	app := fiber.New()

	app.Get("/", controller.Controller)
	app.Get("/albums", controllers.GetAlbums)
	app.Get("/albums/:id", controllers.GetAlbumsById)
	app.Post("/createuser", controllers.CreateUser)
	app.Post("/createproduct", controllers.CreateProduct)
	app.Get("/allusers", controllers.AllUsers)
	app.Get("/allusers/:id", controllers.GetUserById)
	app.Put("/allusers/:id", controllers.UpdateUser)
	log.Fatal(app.Listen(":3000"))

}
