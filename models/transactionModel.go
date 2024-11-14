package models

type Transaction struct {
	ID             int      `json:"id,omitempty"`
	EventID        int      `json:"event_id,omitempty" validate:"required,gt=0"`
	Status         string   `json:"status,omitempty"`
	Price          float64  `json:"price,omitempty" validate:"required"`
	CustomerID     int      `json:"customer_id,omitempty" validate:"required"`
	NumberOfTicket int      `json:"number_of_ticket,omitempty" validate:"gte=0"`
	Customer       Customer `json:"customer"`
}
