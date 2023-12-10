package main

import (
	"log"

	"github.com/adarsh2858/fiber-crm/pkg/handlers"
	"github.com/adarsh2858/fiber-crm/pkg/models"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/crm/users", handlers.GetUsers)
	app.Get("/crm/user/:id", handlers.GetUser)
	app.Post("/crm/user", handlers.AddUser)
	app.Put("/crm/user/:id", handlers.UpdateUser)
	app.Delete("/crm/user/:id", handlers.RemoveUser)
}

func main() {
	app := fiber.New()

	models.InitDatabase()
	setupRoutes(app)

	log.Fatal(app.Listen(3000))
	defer models.DbConn.Close()
}
