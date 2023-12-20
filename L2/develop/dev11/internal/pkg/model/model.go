package model

type Event struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
