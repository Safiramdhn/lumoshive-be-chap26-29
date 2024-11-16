package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
)

type DestinationService struct {
	DestinationRepo repositories.DestinationRepository
}

func NewDestinationService(repo repositories.DestinationRepository) *DestinationService {
	return &DestinationService{repo}
}

func (s *DestinationService) GetDestinationByEventId(event models.Event) (*models.Destination, error) {
	return s.DestinationRepo.GetByEvent(event)
}
