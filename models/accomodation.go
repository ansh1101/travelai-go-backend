package models

type Accommodation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Price    int    `json:"price"`
}
