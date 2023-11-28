package main

import (
	"learnfiber/handlers"
	"learnfiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	handler := handlers.NewHandler()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Get("/", handler.Controller)
	app.Get("/albums", handlers.GetAlbums)
	app.Get("/albums/:id", handlers.GetAlbumsById)
	app.Post("/createproduct", handlers.HandleCreateProduct)
	app.Get("/allusers", handlers.HandleAllUsers)
	app.Get("/allusers/:id", handlers.HandleGetUserById)
	app.Put("/allusers/:id", handlers.HandleUpdateUser)
	app.Post("/user/role/:id", handlers.HandleSetRole)
	app.Post("/register", handlers.HandleRegister)
	app.Get("/allProducts", handlers.HandleAllProducts)
	app.Get("/allproducts/:id", handlers.HandleGetProductById)
	app.Delete("/deleteproduct/:userId/:id", handlers.HandleDeleteProduct)
	app.Get("/allorders", handlers.HandleGetAllOrders)
	app.Get("/orders/user/:user_id", handlers.HandleGetOrdersByUser)
	app.Post("/createorder", handlers.HandleCreateOrder)
	app.Post("/deleteorder/:id", handlers.HandleDeleteOrder)
	app.Post("/login", handlers.HandleLogin)
	app.Get("/currentuser", handlers.HandleCurrentUser)
	app.Get("/logout", handlers.HandleLogout)
	log.Fatal(app.Listen(":3000"))
}

func main() {

	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

}
