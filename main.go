package main

import (
	"testcontact/database"
	"testcontact/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Engine Init
	engine := html.New("views", ".html")

	// Fiber App init
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Database Init
	dsn := "host=localhost user=postgres password=postgres port=5432"
	database.DBConn, _ = database.DatabaseInit(dsn)
	database.MigrateDatabase(database.DBConn, true)

	// Routes Init
	routes.RoutesInit(app)

	app.Listen(":3000")
}
