package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Dbconnect export for global
	Dbconnect *sql.DB
)

// Connect to database exported to main
func Connect() error {

	var err error

	if os.Getenv("POSTGRES_URL") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			fmt.Println(err, "POSTGRES_URL not provided")
		}
	}

	Dbconnect, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	err = Dbconnect.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return nil
}
