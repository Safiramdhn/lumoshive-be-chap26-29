package models

type Review struct {
	ID            int     `json:"id,omitempty"`
	TransactionID int     `json:"transaction_id,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	Comment       string  `json:"comment,omitempty"`
}
