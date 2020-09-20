package models

import (
	"database/sql"
	"encoding/json"
)

// DayRoster ...
type DayRoster struct {
	ID            int    `json:"id"`
	Date          string `json:"date"`
	UpperStaff    string `json:"upperStaff,omitempty"`
	UpperTime     string `json:"upperTime,omitempty"`
	LowerStaff    string `json:"lowerStaff,omitempty"`
	LowerTime     string `json:"lowerTime,omitempty"`
	CustomMessage string `json:"customMessage,omitempty"`
}

// Product ...

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

// https://blog.auberginesolutions.com/how-i-handled-possible-null-values-from-database-rows-in-golang/

// NullString ...
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
