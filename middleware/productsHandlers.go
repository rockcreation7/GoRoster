package middleware

import (
	// package to encode and decode the json into struct and vice versa
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"log" // used to access the request and response object of the api

	"roster-api/db"
	"roster-api/models"

	// package used to covert string into int type
	// used to get the params from the route

	"github.com/gofiber/fiber"
)

type updateProductResponse struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func productTableName() string {
	tableName := os.Getenv("PRODUCT_TABLE")
	if tableName == "" {
		tableName = "Products"
	}
	return tableName
}

// GetAllProduct ...
func GetAllProduct(c *fiber.Ctx) {

	products, err := getAllProducts()
	if err != nil {
		log.Fatalf("Unable to get all product. %v", err)
	}

	fmt.Println(err)
	c.JSON(products)
}

// CreateProduct ...
func CreateProduct(c *fiber.Ctx) {

	Product := new(models.Product)
	// Parse body into struct
	if err := c.BodyParser(Product); err != nil {
		c.Status(400).Send(err)
		return
	}

	// call insert product function and pass the product
	insertID, err := insertProduct(Product)

	if err != nil {
		c.JSON(err)
		return
	}

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Product created successfully",
	}

	// send the response
	c.JSON(res)
}

// UpdateProduct update product's detail in the postgres db
func UpdateProduct(c *fiber.Ctx) {

	Product := new(models.Product)
	// Parse body into struct
	if err := c.BodyParser(Product); err != nil {
		c.Status(400).Send(err)
		return
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	updatedRows := updateProduct(id, Product)

	// format the message string
	msg := fmt.Sprintf("Product updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := updateProductResponse{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(res)
}

// DeleteProduct delete product's detail in the postgres db
func DeleteProduct(c *fiber.Ctx) {

	// call the deleteProduct

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteProduct(id)

	// format the message string
	msg := fmt.Sprintf("Product updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := updateProductResponse{
		ID:      id,
		Message: msg,
	}

	c.JSON(res)
}

// GetProduct get product by date
func GetProduct(c *fiber.Ctx) {
	// convert the id type from string to int

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	product, err := getProduct(id)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	c.JSON(product)
}

//------------------------- handler functions ----------------

func insertProduct(Product *models.Product) (int64, error) {

	db := db.Dbconnect

	sqlStatement := `INSERT INTO ` + productTableName() + ` (Name, Price, Cost, Qty, Code) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, Product.Name, Product.Price, Product.Cost, Product.Qty, Product.Code).Scan(&id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		fmt.Printf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, err
}

func getAllProducts() ([]models.Product, error) {
	// create the postgres db connection
	db := db.Dbconnect
	var products []models.Product
	sqlStatement := `SELECT * from ` + productTableName()
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Cost,
			&product.Qty,
			&product.Code,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		products = append(products, product)
	}

	return products, err
}

// update product in the DB
func updateProduct(id int, product *models.Product) int64 {

	// create the postgres db connection
	db := db.Dbconnect

	// create the update sql query //date
	sqlStatement := `UPDATE ` + productTableName() + ` SET Name=$1, Price=$2, Cost=$3, Qty=$4, Code=$5 WHERE id=$6`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, product.Name, product.Price, product.Cost, product.Qty, product.Code, product.ID)

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

// delete product in the DB
func deleteProduct(id int) int64 {

	// create the postgres db connection
	db := db.Dbconnect

	// create the delete sql query
	sqlStatement := `DELETE FROM ` + productTableName() + ` WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

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
func getProduct(id int) (models.Product, error) {
	// create the postgres db connection
	db := db.Dbconnect

	// create a user of models.User type
	var product models.Product

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + productTableName()

	// execute the sql statement
	row := db.QueryRow(sqlStatement)

	// unmarshal the row object to user
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Cost,
		&product.Qty,
		&product.Code,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Printf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return product, err
}
