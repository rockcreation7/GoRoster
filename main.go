package main

import (
	"fmt"
	"log"
	"os"
	"roster-api/db"
	"roster-api/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func main() {

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Dbconnect.Close()
	fmt.Println("db connection success")

	app := fiber.New()
	app.Use(cors.New())

	router.Route(app)

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}

	app.Listen(Port)
}
