package router

import (
	"fmt"
	"roster-api/middleware"

	"github.com/gofiber/fiber/v2"
)

// Route export to main
func Route(app *fiber.App) {

	api := app.Group("/api/roster")
	productAPI := app.Group("/api/product")

	api.Get("/get", middleware.GetAllRoster)
	api.Get("/get/:date", middleware.GetRoster)
	api.Post("/insert", middleware.CreateRoster)
	api.Put("/update/:date", middleware.UpdateRoster)
	api.Delete("/delete/:date", middleware.DeleteRoster)

	productAPI.Get("/get", middleware.GetAllProduct)
	productAPI.Post("/insert", middleware.CreateProduct)
	productAPI.Put("/update/:id", middleware.UpdateProduct)
	productAPI.Get("/get/:id", middleware.GetProduct)
	productAPI.Delete("/delete/:id", middleware.DeleteProduct)

	app.Post("/image", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		// Save file to root directory:
		return c.SaveFile(file, fmt.Sprintf("./dist/img/%s", file.Filename))
	})

	app.Static("/", "dist")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})

}
