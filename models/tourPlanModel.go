package models

type TourPlan struct {
	ID          int      `json:"id,omitempty"`
	EventID     int      `json:"event_id,omitempty"`
	Activities  []string `json:"activities,omitempty"`
	Description string   `json"description,omitempty"`
}
