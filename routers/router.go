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
	log := config.IntiLogger()
	validator := config.InitValidator()

	tourRepo := repositories.NewRepository[models.TourData](db, log)
	tourService := services.NewTourService(tourRepo)
	tourHandler := handlers.NewTourHandler(tourService, log, validator)

	customerRepo := repositories.NewCustomerRepo(db, log)
	customerService := services.NewCustomerService(*customerRepo)

	transactionRepo := repositories.NewTransactionRepo(db, log)
	transactionService := services.NewTransactionService(*transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, customerService, log, validator)

	r.Route("/api", func(r chi.Router) {
		r.Route("/tour", func(r chi.Router) {
			r.Get("/", tourHandler.GetTourDataHandler)
			r.Get("/{id}", tourHandler.GetTourDetailsHandler)
		})

		r.Post("/bookings", transactionHandler.CreateTransactionHandler)
	})

	return r
}
