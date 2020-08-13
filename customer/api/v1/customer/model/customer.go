package model

// Customer model
type Customer struct {
    ID int `json:"id" db:"id"`
	Code string `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
}