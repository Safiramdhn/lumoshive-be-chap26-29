package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
)

type TourPlanService struct {
	TourPlanRepo repositories.TourPlanRepository
}

func NewTourPlanService(repo repositories.TourPlanRepository) *TourPlanService {
	return &TourPlanService{repo}
}

func (t *TourPlanService) GetTourPlanByEventId(event models.Event) ([]models.TourPlan, error) {
	return t.TourPlanRepo.GetByEvent(event)
}
