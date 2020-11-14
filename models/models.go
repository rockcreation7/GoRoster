package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// DayRoster - format for database
type DayRoster struct {
	ID            int        `json:"id"`
	Date          CustomTime `json:"date"`
	UpperStaff    string     `json:"upperStaff,omitempty"`
	UpperTime     string     `json:"upperTime,omitempty"`
	LowerStaff    string     `json:"lowerStaff,omitempty"`
	LowerTime     string     `json:"lowerTime,omitempty"`
	CustomMessage string     `json:"customMessage,omitempty"`
}

// CustomTime - defined
type CustomTime time.Time

const ctLayout = "2006-01-02 15:04:05 Z07:00"

// UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	cs := s + ` 00:00:00 Z`
	fmt.Println(cs, "CustomTime")
	nt, err := time.Parse(ctLayout, cs)
	fmt.Println(nt, "NewTime", err)
	*ct = CustomTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time().String()), nil
}

// Time returns the time in the custom format
func (ct *CustomTime) Time() time.Time {
	t := time.Time(*ct)
	return t
}

// Make it as string
/*
func (ct *CustomTime) Time() time.Time {
	t := time.Time(*ct)
	return t
}
*/

// Product model for cashier
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Qty      int     `json:"qty,string,omitempty"`
	Price    float64 `json:"price,string,omitempty"`
	Cost     int     `json:"cost,string,omitempty"`
	Code     int     `json:"code,string,omitempty"`
	Catagory string  `json:"catagory,omitempty"`
	Imgurl   string  `json:"Imgurl,omitempty"`
}

// NullString ...
// https://blog.auberginesolutions.com/how-i-handled-possible-null-values-from-database-rows-in-golang/
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}
