package main

import (
	"learnfiber/database"
	"learnfiber/handlers"
	"learnfiber/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	handler := handlers.NewHandler()

	protected := app.Group("/protected", middleware.CheckAuth)

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
	protected.Get("/allusers", handlers.HandleAllUsers)
	protected.Get("/user/:id", handlers.HandleGetUserById)
	protected.Put("/user/:id", handlers.HandleUpdateUser)
	app.Post("/user/role/:id", handlers.HandleSetRole)
	app.Post("/register", handlers.HandleRegister)
	protected.Get("/allProducts", handlers.HandleAllProducts)
	app.Get("/allproducts/:id", handlers.HandleGetProductById)
	protected.Delete("/deleteproduct/:userId/:id", handlers.HandleDeleteProduct)
	protected.Get("/allorders", handlers.HandleGetAllOrders)
	app.Get("/orders/user/:user_id", handlers.HandleGetOrdersByUser)
	app.Post("/createorder", handlers.HandleCreateOrder)
	protected.Post("/deleteorder/:id", handlers.HandleDeleteOrder)
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
