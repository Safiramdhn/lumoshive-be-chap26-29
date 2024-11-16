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

type DestinationHandler struct {
	DestinationService services.DestinationService
	Logger             *zap.Logger
	Validator          *validator.Validate
}

func NewDestinationHandler(destinationService services.DestinationService, log *zap.Logger, validator *validator.Validate) *DestinationHandler {
	return &DestinationHandler{DestinationService: destinationService, Logger: log, Validator: validator}
}

func (h *DestinationHandler) GetDestinationLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.Logger.Error("Invalid request method", zap.String("method", r.Method), zap.String("handler", "DestinationHandler"), zap.String("function", "GetDestinationLocationHandler"))
		JsonResp.SendError(w, http.StatusMethodNotAllowed, r.Method)
	}

	idParam := chi.URLParam(r, "id")
	eventId, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Error("Error parsing event ID "+err.Error(), zap.String("handler", "DestinatonHandler"), zap.String("function", "GetDestinationLocationHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Invalid event ID", err.Error())
		return
	}

	var event models.Event
	event.ID = eventId
	err = h.Validator.Struct(event)
	if err != nil {
		h.Logger.Error("Validation error for event "+err.Error(), zap.String("handler", "DestinatonHandler"), zap.String("function", "GetDestinationLocationHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	destination, err := h.DestinationService.GetDestinationByEventId(event)
	if err != nil {
		h.Logger.Error("Error getting location "+err.Error(), zap.String("handler", "DestinationHandler"), zap.String("function", "GetDestinationLocationHandler"))
	}

	JsonResp.SendSuccess(w, destination, "Location retrieved successfully")
}
