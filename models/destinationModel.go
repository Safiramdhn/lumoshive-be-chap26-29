package models

type Destination struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Location    string  `json:"location,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Map         string  `json:"map,omitempty"`
}
