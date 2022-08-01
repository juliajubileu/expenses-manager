package models

type Expense struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Amount string `json:"amount"`
	Date string `json:"date"`
}

var Expenses []Expense
