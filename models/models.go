package models

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
