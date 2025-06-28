package model

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}
