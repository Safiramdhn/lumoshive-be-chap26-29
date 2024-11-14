package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID            int            `json:"id,omitempty" validate:"gt=0"`
	DestinationID int            `json:"destination_id,omitempty"`
	Date          time.Time      `json:"date,omitempty"`
	Title         string         `json:"title,omitempty"`
	Description   string         `json:"description,omitempty"`
	PhotoUrl      sql.NullString `json:"photo_url,omitempty"`
	Price         float64        `json:"price,omitempty"`
}
