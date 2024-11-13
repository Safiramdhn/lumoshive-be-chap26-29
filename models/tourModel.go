package models

type TourData struct {
	Transaction       Transaction
	Event             Event
	Destination       Destination
	AverageRating     float64
	TotalTransactions int
}
