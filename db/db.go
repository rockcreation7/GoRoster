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

	if !checkIfTableExist("Products", Dbconnect) {
		migrateProduct(Dbconnect)
	}

	if !checkIfTableExist("Rosters", Dbconnect) {
		migrateDayRoster(Dbconnect)
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

func migrateProduct(Dbconnect *sql.DB) {
	stmt, err := Dbconnect.Prepare("CREATE Table products(id SERIAL PRIMARY KEY, cost int, Imgurl varchar(50), catagory varchar(50),code int, price int, qty varchar(50), name varchar(50))")

	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Product table created successfully")
	}
}

func migrateDayRoster(Dbconnect *sql.DB) {
	stmt, err := Dbconnect.Prepare("CREATE Table Rosters(id SERIAL PRIMARY KEY, Date DATE, UpperStaff varchar(50), UpperTime varchar(255), LowerStaff varchar(50), LowerTime varchar(255), CustomMessage varchar(50))")

	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Product table created successfully")
	}
}

// stmt, err = Dbconnect.Prepare("ALTER TABLE products DROP email;")
func checkIfTableExist(tableName string, Dbconnect *sql.DB) bool {
	_, tableCheck := Dbconnect.Query("select * from " + tableName + ";")
	if tableCheck == nil {
		// fmt.Println("table is there", res)
		return true
	}
	// fmt.Println("table not there", res)
	return false
}
