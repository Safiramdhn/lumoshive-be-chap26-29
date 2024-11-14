package models

type Customer struct {
	ID    int    `json:"id" validate:"gt = 0"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,phone"`
}
