package models

type Transaction struct {
	ID      int     `json:"id,omitempty"`
	EventID int     `json:"event_id,omitempty"`
	Status  string  `json:"status,omitempty"`
	Price   float64 `json:"price,omitempty"`
}
