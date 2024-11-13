package handlers

import (
	"golang-beginner-chap28/services"
	"golang-beginner-chap28/utils"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type TourHandler struct {
	TourService services.TourService
	Logger      *zap.Logger
}

func NewTourHandler(tourService services.TourService, log *zap.Logger) *TourHandler {
	return &TourHandler{TourService: tourService, Logger: log}
}

var JsonResp = &utils.JSONResponse{}

const (
	DateFormat      = "2006-01-02"
	DefaultPage     = 1
	DefaultPageSize = 10
)

func parseDateFilter(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil // Return zero value if date is empty
	}
	return time.Parse(DateFormat, dateStr)
}

func parseIntQueryParam(param string, defaultVal int) int {
	value, err := strconv.Atoi(param)
	if err != nil || value < 1 {
		return defaultVal
	}
	return value
}

func (h *TourHandler) GetTourDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.Logger.Error("Method not allowed"+r.Method, zap.String("handler", "TourHandler"), zap.String("function", "GetTourDataHandler"))
		JsonResp.SendError(w, http.StatusMethodNotAllowed, "Method not allowed "+r.Method)
		return
	}

	// Handler
	dateFilter := r.URL.Query().Get("date")
	formatedDateFilter, err := parseDateFilter(dateFilter)
	if err != nil {
		h.Logger.Error("Error parsing date filter "+err.Error(), zap.String("handler", "TourHandler"), zap.String("function", "GetTourDataHandler"))
		JsonResp.SendError(w, http.StatusBadRequest, "Invalid date format. Expected format: YYYY-MM-DD", err.Error())
		return
	}

	priceSort := r.URL.Query().Get("sort_by_price")
	page := parseIntQueryParam(r.URL.Query().Get("page"), DefaultPage)
	pageSize := parseIntQueryParam(r.URL.Query().Get("page_size"), DefaultPageSize)

	h.Logger.Info("Get tour data", zap.String("handler", "TourHandler"), zap.String("function", "GetTourDataHandler"))
	tours, totalItem, err := h.TourService.GetTourData(formatedDateFilter, priceSort, page, pageSize)
	if err != nil {
		h.Logger.Error("Error fetching tours data "+err.Error(), zap.String("handler", "TourHandler"), zap.String("function", "GetTourDataHandler"))
		JsonResp.SendError(w, http.StatusInternalServerError, "Error fetching tours data", err.Error())
		return
	}
	totalPage := totalItem / pageSize
	JsonResp.SendPaginatedResponse(w, tours, page, pageSize, totalItem, totalPage, "Tour data successfully fetched")
}
