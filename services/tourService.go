package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
	"time"
)

type TourService struct {
	repository repositories.Repository[models.TourData]
}

func NewTourService(repository repositories.Repository[models.TourData]) *TourService {
	return &TourService{repository}
}

func (s *TourService) GetTourData(dateFilter time.Time, sortByPrice string, page, pageSize int) ([]models.TourData, int, error) {
	return s.repository.GetTourData(dateFilter, sortByPrice, page, pageSize)
}
