package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
)

type EventService struct {
	EventRepo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) *EventService {
	return &EventService{repo}
}

func (s *EventService) GetEventById(id int) (*models.Event, error) {
	return s.EventRepo.GetById(id)
}
