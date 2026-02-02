package dto

import (
	"time"

	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

type PatientResponse struct {
	ID        string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Birthdate time.Time      `json:"birth_date"`
	Sex       entitiy.Gender `json:"sex"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type PatientRequest struct {
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Birthdate time.Time      `json:"birth_date"`
	Sex       entitiy.Gender `json:"sex"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
}
