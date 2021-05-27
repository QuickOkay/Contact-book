package routes

import (
	"testcontact/handlers"

	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {
	// Page Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// API Routes
	api := app.Group("/api")
	api.Post("/add-contact", handlers.AddContact)
	api.Post("/remove-contact", handlers.RemoveContact)
	api.Get("/get-contacts", handlers.GetContacts)
	api.Post("/edit-contact", handlers.EditContact)
}
