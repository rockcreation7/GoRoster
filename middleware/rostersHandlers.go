package middleware

import (
	// package to encode and decode the json into struct and vice versa
	"database/sql"
	"fmt"
	"reflect"

	// used to access the request and response object of the api
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

// DayRosterResponse - response format
type DayRosterResponse struct {
	ID            int       `json:"id"`
	Date          time.Time `json:"date"`
	UpperStaff    string    `json:"upperStaff"`
	UpperTime     string    `json:"upperTime"`
	LowerStaff    string    `json:"lowerStaff"`
	LowerTime     string    `json:"lowerTime"`
	CustomMessage string    `json:"customMessage"`
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
		panic("Unable to get all roster . %v")
	}

	/* 	sort.SliceStable(rosters, func(i, j int) bool {
		fmt.Println(rosters[i].Date.String(), "sort")
		return rosters[i].Date.String() < rosters[j].Date.String()
	}) */

	var dayRosterResponse []DayRosterResponse

	for _, v := range rosters {
		var ResponseObject DayRosterResponse
		ResponseObject.Date = v.Date.Time()
		ResponseObject.ID = v.ID
		ResponseObject.UpperStaff = v.UpperStaff
		ResponseObject.UpperTime = v.UpperTime
		dayRosterResponse = append(dayRosterResponse, ResponseObject)
	}

	return c.JSON(dayRosterResponse)
}

/*
// CustomBodyParser - CustomBodyParser
func (c *Ctx) CustomBodyParser(out interface{}) error {

	fmt.println("here")
	// Get decoder from pool
	schemaDecoder := decoderPool.Get().(*schema.Decoder)

	defer decoderPool.Put(schemaDecoder)

	// Get content-type
	ctype := getString(c.fasthttp.Request.Header.ContentType())

	// Parse body accordingly
	if strings.HasPrefix(ctype, MIMEApplicationJSON) {
		schemaDecoder.SetAliasTag("json")
		return json.Unmarshal(c.fasthttp.Request.Body(), out)
	} else if strings.HasPrefix(ctype, MIMEApplicationForm) {
		schemaDecoder.SetAliasTag("form")
		data := make(map[string][]string)
		c.fasthttp.PostArgs().VisitAll(func(key []byte, val []byte) {
			data[getString(key)] = append(data[getString(key)], getString(val))
		})
		return schemaDecoder.Decode(out, data)
	} else if strings.HasPrefix(ctype, MIMEMultipartForm) {
		schemaDecoder.SetAliasTag("form")
		data, err := c.fasthttp.MultipartForm()
		if err != nil {
			return err
		}
		return schemaDecoder.Decode(out, data.Value)
	} else if strings.HasPrefix(ctype, MIMETextXML) || strings.HasPrefix(ctype, MIMEApplicationXML) {
		schemaDecoder.SetAliasTag("xml")
		return xml.Unmarshal(c.fasthttp.Request.Body(), out)
	}
	// No suitable content type found
	return fmt.Errorf("bodyparser: cannot parse content-type: %v", ctype)
}
*/
var timeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("2006-02-01", value); err == nil {
		fmt.Println(v, "v  err")
		return reflect.ValueOf(v)
	}
	return reflect.Value{} // this is the same as the private const invalidType
}

// CreateRoster - insertRoster
func CreateRoster(c *fiber.Ctx) error {

	confCustomTime := fiber.CustomRegister{
		models.CustomTime{},
		timeConverter,
	}

	confCustomTime2 := fiber.CustomRegister{
		models.CustomTime2{},
		timeConverter,
	}

	confs := []fiber.CustomRegister{confCustomTime2, confCustomTime}

	Roster := new(models.DayRoster)
	// Parse body into struct
	if err := c.CustomFormParser(Roster, confs); err != nil {
		panic(err)
	}

	// call insert roster function and pass the roster
	insertID, err := insertRoster(Roster)

	if err != nil {
		panic("error on insert roster")
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
		panic("error Parse body into struct")
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

	var dayRosterResponse DayRosterResponse

	// https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
	// https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go

	dayRosterResponse.ID = roster.ID
	dayRosterResponse.Date = roster.Date.Time()
	dayRosterResponse.UpperStaff = roster.UpperTime
	dayRosterResponse.UpperTime = roster.UpperTime
	dayRosterResponse.CustomMessage = roster.CustomMessage

	if err != nil {
		panic("get Roster err")
	}

	return c.JSON(dayRosterResponse)
}

//------------------------- handler functions ----------------

func insertRoster(Roster *models.DayRoster) (int64, error) {

	db := db.Dbconnect

	sqlStatement := `INSERT INTO ` + tableName() + ` (Date, UpperStaff, UpperTime, LowerStaff, LowerTime, CustomMessage) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, Roster.Date.String(), Roster.UpperStaff, Roster.UpperTime, Roster.LowerStaff, Roster.LowerTime, Roster.CustomMessage).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, err
}

func getAllRosters() ([]models.DayRoster, error) {
	// create the postgres db connection
	db := db.Dbconnect

	today := time.Now()
	sevenDayLater := today.AddDate(0, 0, 14)

	fmt.Println(today)
	fmt.Println(sevenDayLater)

	// SELECT * from Rosters where (date <= '2020-09-09' AND date >= '2020-09-04')
	var rosters []models.DayRoster
	sqlStatement := `SELECT * from ` + tableName() + ` where (date <= $1 AND date >= $2)`
	rows, err := db.Query(sqlStatement, sevenDayLater, today)
	if err != nil {
		panic("Unable to execute the query. %v")
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
			panic("Unable to scan the row. %v")
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
		panic(err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		panic(err)
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
		panic("Unable to execute the query. %v")
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		panic("Error while checking the affected rows. %v")
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
		panic("No rows were returned!")
	case nil:
		return roster, nil
	default:
		panic("Unable to scan the row. %v")
	}

}
