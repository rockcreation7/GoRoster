package router

import (
	"roster-api/middleware"

	"github.com/gofiber/fiber"
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

	app.Static("/", "dist")
	app.Get("/*", func(c *fiber.Ctx) {
		c.SendFile("dist/index.html")
	})
}
