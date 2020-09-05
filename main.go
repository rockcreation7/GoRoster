package main

import (
	"roster-api/middleware"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api/roster")

	app.Static("/", "./dist")

	api.Get("/get", middleware.GetAllRoster)
	api.Get("/get/:date", middleware.GetRoster)
	api.Post("/insert", middleware.CreateRoster)
	api.Put("/update/:date", middleware.UpdateRoster)
	api.Delete("/delete/:date", middleware.DeleteRoster)

	app.Listen(3000)
}
