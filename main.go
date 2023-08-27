package main

import (
	"learnfiber/controllers"
	"learnfiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	controller := controllers.NewHandler()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Get("/", controller.Controller)
	app.Get("/albums", controllers.GetAlbums)
	app.Get("/albums/:id", controllers.GetAlbumsById)
	app.Post("/createproduct", controllers.CreateProduct)
	app.Get("/allusers", controllers.AllUsers)
	app.Get("/allusers/:id", controllers.GetUserById)
	app.Put("/allusers/:id", controllers.UpdateUser)
	app.Post("/register", controllers.Register)
	app.Get("/allProducts", controllers.AllProducts)
	app.Get("/allproducts/:id", controllers.GetProductById)
	app.Get("/allorders", controllers.AllOrders)
	app.Post("/createorder", controllers.CreateOrder)
	app.Post("/login", controllers.Login)
	app.Get("/currentuser", controllers.CurrentUser)
	app.Get("/logout", controllers.Logout)
	log.Fatal(app.Listen(":3000"))
}

func main() {

	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

}
