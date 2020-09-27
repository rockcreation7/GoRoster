package middleware

import (
	// package to encode and decode the json into struct and vice versa
	"database/sql"
	"fmt"

	"log" // used to access the request and response object of the api
	"os"
	"time"

	"roster-api/db"
	"roster-api/models"

	// package used to covert string into int type
	// used to get the params from the route

	"github.com/gofiber/fiber/v2"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type updateResponse struct {
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
}

func tableName() string {
	tableName := os.Getenv("TABLE")
	if tableName == "" {
		tableName = "Rosters"
	}
	return tableName
}

// GetAllRoster ...
func GetAllRoster(c *fiber.Ctx) error {

	rosters, err := getAllRosters()
	if err != nil {
		log.Fatalf("Unable to get all roster. %v", err)
	}

	fmt.Println(err)
	return c.JSON(rosters)
}

// CreateRoster ...
func CreateRoster(c *fiber.Ctx) error {

	Roster := new(models.DayRoster)
	// Parse body into struct
	if err := c.BodyParser(Roster); err != nil {
		return err
	}

	// call insert roster function and pass the roster
	insertID, err := insertRoster(Roster)

	if err != nil {
		return c.JSON(err)
	}

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Roster created successfully",
	}

	// send the response
	return c.JSON(res)
}

// UpdateRoster update roster's detail in the postgres db
func UpdateRoster(c *fiber.Ctx) error {

	Roster := new(models.DayRoster)
	// Parse body into struct
	if err := c.BodyParser(Roster); err != nil {
		return err
	}
	date := c.Params("date")
	updatedRows := updateRoster(date, Roster)
	// error on invalidate input
	// format the message string
	msg := fmt.Sprintf("Roster updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := updateResponse{
		Date:    date,
		Message: msg,
	}

	// send the response
	return c.JSON(res)
}

// DeleteRoster delete roster's detail in the postgres db
func DeleteRoster(c *fiber.Ctx) error {

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

	return c.JSON(res)
}

// GetRoster get roster by date
func GetRoster(c *fiber.Ctx) error {

	date := c.Params("date")
	// convert the id type from string to int

	roster, err := getRoster(date)

	if err != nil {
		return err
	}

	return c.JSON(roster)
}

//------------------------- handler functions ----------------

func insertRoster(Roster *models.DayRoster) (int64, error) {

	db := db.Dbconnect

	sqlStatement := `INSERT INTO ` + tableName() + ` (Date, UpperStaff, UpperTime, LowerStaff, LowerTime, CustomMessage) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

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
	db := db.Dbconnect

	today := time.Now()
	sevenDayLater := today.AddDate(0, 0, 7)

	fmt.Println(today)
	fmt.Println(sevenDayLater)

	// SELECT * from Rosters where (date <= '2020-09-09' AND date >= '2020-09-04')
	var rosters []models.DayRoster
	sqlStatement := `SELECT * from ` + tableName() + ` where (date <= $1 AND date >= $2)`
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
	db := db.Dbconnect

	// create the update sql query //date
	sqlStatement := `UPDATE ` + tableName() + ` SET UpperStaff=$2, LowerStaff=$3, UpperTime=$4, LowerTime=$5, CustomMessage=$6 WHERE date=$1`

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
	db := db.Dbconnect

	// create the delete sql query
	sqlStatement := `DELETE FROM ` + tableName() + ` WHERE date=$1`

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
	db := db.Dbconnect

	// create a user of models.User type
	var roster models.DayRoster

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + tableName() + ` WHERE date=$1`

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
