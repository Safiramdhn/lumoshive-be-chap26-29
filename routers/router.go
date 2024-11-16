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

	tourPlanRepo := repositories.NewTourPlanRepository(db, log)
	tourPlanService := services.NewTourPlanService(*tourPlanRepo)
	tourPlanHandler := handlers.NewTourPlanHandler(tourPlanService, log, validator)

	destinationRepo := repositories.NewDestinationRepository(db, log)
	destinationService := services.NewDestinationService(*destinationRepo)
	destinationHandler := handlers.NewDestinationHandler(*destinationService, log, validator)

	r.Route("/api", func(r chi.Router) {
		r.Route("/tour", func(r chi.Router) {
			r.Get("/", tourHandler.GetTourDataHandler)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", tourHandler.GetTourDetailsHandler)
				r.Get("/tour-plan", tourPlanHandler.GetTourPlanHandler)
				r.Get("/location", destinationHandler.GetDestinationLocationHandler)
			})
		})

		r.Post("/bookings", transactionHandler.CreateTransactionHandler)
	})

	return r
}
