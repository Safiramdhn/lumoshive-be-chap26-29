package handlers

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type TourPlanHandler struct {
	TourPlanService *services.TourPlanService
	Logger          *zap.Logger
	Validator       *validator.Validate
}

func NewTourPlanHandler(tourPlanService *services.TourPlanService, log *zap.Logger, validator *validator.Validate) *TourPlanHandler {
	return &TourPlanHandler{TourPlanService: tourPlanService, Logger: log, Validator: validator}
}

func (htp *TourPlanHandler) GetTourPlanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		htp.Logger.Error("invalid request method", zap.String("method", r.Method), zap.String("handler", "TourPlanHandler"), zap.String("function", "GetTourPlanHandler"))
		JsonResp.SendError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}

	idParam := chi.URLParam(r, "id")
	eventId, err := strconv.Atoi(idParam)
	if err != nil {
		htp.Logger.Error("Error parsing event ID "+err.Error(), zap.String("handler", "TourPlanHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Invalid event ID", err.Error())
		return
	}

	event := models.Event{ID: eventId}
	err = htp.Validator.Struct(event)
	if err != nil {
		htp.Logger.Error("Error validating event "+err.Error(), zap.String("handler", "TourPlanHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Validation error", err.Error())
		return
	}
	tourPlans, err := htp.TourPlanService.GetTourPlanByEventId(event)
	if err != nil {
		htp.Logger.Error("Error getting tour plans "+err.Error(), zap.String("handler", "TourPlanHandler"))
		JsonResp.SendError(w, http.StatusInternalServerError, "Error getting tour plans", err.Error())
		return
	}
	JsonResp.SendSuccess(w, tourPlans, "Tour plans fetched successfully")
}
