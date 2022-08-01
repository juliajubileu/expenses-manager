package models

type Income struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Amount string `json:"amount"`
	Date string `json:"date"`
}

var Incomes []Income
