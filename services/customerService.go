package services

import (
	"golang-beginner-chap28/models"
	"golang-beginner-chap28/repositories"
)

type CustomerService struct {
	CustomerRepo repositories.CustomerRepo
}

func NewCustomerService(customerRepo repositories.CustomerRepo) *CustomerService {
	return &CustomerService{customerRepo}
}

func (cs *CustomerService) CreateCustomer(customerInput *models.Customer) (*models.Customer, error) {
	return cs.CustomerRepo.Create(*customerInput)
}
