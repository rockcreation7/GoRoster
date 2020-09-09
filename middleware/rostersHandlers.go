package middleware

import (
	// package to encode and decode the json into struct and vice versa
	"database/sql"
	"fmt"

	"log" // used to access the request and response object of the api
	"os"
	"time"

	// used to read the environment variable
	"roster-api/models"

	// package used to covert string into int type
	// used to get the params from the route

	"github.com/gofiber/fiber"

	// db driver
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type updateResponse struct {
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetAllRoster ...
func GetAllRoster(c *fiber.Ctx) {
	rosters, err := getAllRosters()
	if err != nil {
		log.Fatalf("Unable to get all roster. %v", err)
	}

	fmt.Println(err)
	c.JSON(rosters)
}

// CreateRoster ...
func CreateRoster(c *fiber.Ctx) {

	Roster := new(models.DayRoster)
	// Parse body into struct
	if err := c.BodyParser(Roster); err != nil {
		c.Status(400).Send(err)
		return
	}

	// call insert roster function and pass the roster
	insertID, err := insertRoster(Roster)

	if err != nil {
		c.JSON(err)
		return
	}

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Roster created successfully",
	}

	// send the response
	c.JSON(res)
}

// UpdateRoster update roster's detail in the postgres db
func UpdateRoster(c *fiber.Ctx) {

	Roster := new(models.DayRoster)
	// Parse body into struct
	if err := c.BodyParser(Roster); err != nil {
		c.Status(400).Send(err)
		return
	}
	date := c.Params("date")
	updatedRows := updateRoster(date, Roster)

	// format the message string
	msg := fmt.Sprintf("Roster updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := updateResponse{
		Date:    date,
		Message: msg,
	}

	// send the response
	c.JSON(res)
}

// DeleteRoster delete roster's detail in the postgres db
func DeleteRoster(c *fiber.Ctx) {

	// call the deleteRoster

	date := c.Params("date")
	deletedRows := deleteRoster(date)

	// format the message string
	msg := fmt.Sprintf("Roster updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := updateResponse{
		Date:    date,
		Message: msg,
	}

	c.JSON(res)
}

// GetRoster get roster by date
func GetRoster(c *fiber.Ctx) {

	date := c.Params("date")
	// convert the id type from string to int

	roster, err := getRoster(date)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	c.JSON(roster)
}

//------------------------- handler functions ----------------

func insertRoster(Roster *models.DayRoster) (int64, error) {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO Rosters (Date, UpperStaff, UpperTime, LowerStaff, LowerTime, CustomMessage) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, Roster.Date, Roster.UpperStaff, Roster.UpperTime, Roster.LowerStaff, Roster.LowerTime, Roster.CustomMessage).Scan(&id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		fmt.Printf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, err
}

func getAllRosters() ([]models.DayRoster, error) {
	// create the postgres db connection
	db := createConnection()
	defer db.Close()

	today := time.Now()
	sevenDayLater := today.AddDate(0, 0, 7)

	fmt.Println(today)
	fmt.Println(sevenDayLater)

	// SELECT * from Rosters where (date <= '2020-09-09' AND date >= '2020-09-04')
	var rosters []models.DayRoster
	sqlStatement := `SELECT * from Rosters where (date <= $1 AND date >= $2)`
	rows, err := db.Query(sqlStatement, sevenDayLater, today)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var roster models.DayRoster
		err = rows.Scan(
			&roster.ID,
			&roster.Date,
			&roster.UpperStaff,
			&roster.UpperTime,
			&roster.LowerStaff,
			&roster.LowerTime,
			&roster.CustomMessage,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		rosters = append(rosters, roster)
	}

	return rosters, err
}

// update roster in the DB
func updateRoster(date string, roster *models.DayRoster) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query //date
	sqlStatement := `UPDATE rosters SET UpperStaff=$2, LowerStaff=$3, UpperTime=$4, LowerTime=$5, CustomMessage=$6 WHERE date=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, date, roster.UpperStaff, roster.LowerStaff, roster.UpperTime, roster.LowerTime, roster.CustomMessage)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete roster in the DB
func deleteRoster(date string) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM rosters WHERE date=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, date)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// get one user from the DB by its userid
func getRoster(date string) (models.DayRoster, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var roster models.DayRoster

	// create the select sql query
	sqlStatement := `SELECT * FROM Rosters WHERE date=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, date)

	// unmarshal the row object to user
	err := row.Scan(
		&roster.ID,
		&roster.Date,
		&roster.UpperStaff,
		&roster.UpperTime,
		&roster.LowerStaff,
		&roster.LowerTime,
		&roster.CustomMessage,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return roster, nil
	case nil:
		return roster, nil
	default:
		log.Printf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return roster, err
}

// GetAllRoster send all Roster
func createConnection() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
