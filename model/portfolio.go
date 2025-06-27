package model

type Portfolio struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Image            string `json:"image"`
	ShortDescription string `json:"short_description"`
	Client           string `json:"client"`
	Website          string `json:"website"`
	LongDescription  string `json:"long_description"`
}
