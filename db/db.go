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
	isTableExist := checkIfTableExist(Dbconnect)
	if !isTableExist {
		migrate(Dbconnect)
	}

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

func migrate(Dbconnect *sql.DB) {
	stmt, err := Dbconnect.Prepare("CREATE Table products(id int NOT NULL, cost int, Imgurl varchar(50), catagory varchar(50),code int, price int, qty varchar(50), name varchar(50), PRIMARY KEY (id))")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully")
	}
}

// stmt, err = Dbconnect.Prepare("ALTER TABLE products DROP email;")
func checkIfTableExist(Dbconnect *sql.DB) bool {
	_, tableCheck := Dbconnect.Query("select * from products;")
	if tableCheck == nil {
		// fmt.Println("table is there", res)
		return true
	}
	// fmt.Println("table not there", res)
	return false
}
