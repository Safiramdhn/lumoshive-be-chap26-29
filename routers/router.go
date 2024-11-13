package routers

import (
	"golang-beginner-chap28/config"
	"golang-beginner-chap28/handlers"
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
	"golang-beginner-chap28/services"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	db := config.NewPostgresDB()

	tourRepo := repositories.NewRepository[models.TourData](db)
	tourService := services.NewTourService(tourRepo)
	tourHandler := handlers.NewTourHandler(*tourService)

	r.Get("/events", tourHandler.GetTourDataHandler)

	return r
}
