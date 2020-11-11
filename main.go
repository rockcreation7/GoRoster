package main

import (
	"fmt"
	"log"
	"os"
	"roster-api/db"
	"roster-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
)

func main() {

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Dbconnect.Close()

	fmt.Println("db connection success")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New())

	router.Route(app)

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	log.Fatal(app.Listen(":" + Port))
}

// live-reload
// https://techinscribed.com/5-ways-to-live-reloading-go-applications/
