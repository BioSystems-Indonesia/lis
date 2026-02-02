package dto

import (
	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

// ToEntity converts PatientRequest to Patient entity
func (req *PatientRequest) ToEntity(id string) *entitiy.Patient {
	return &entitiy.Patient{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthdate: req.Birthdate,
		Sex:       req.Sex,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
	}
}

// ToPatientResponse converts Patient entity to PatientResponse
func ToPatientResponse(patient *entitiy.Patient) *PatientResponse {
	if patient == nil {
		return nil
	}

	return &PatientResponse{
		ID:        patient.ID,
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		Birthdate: patient.Birthdate,
		Sex:       patient.Sex,
		Address:   patient.Address,
		Phone:     patient.Phone,
		Email:     patient.Email,
		CreatedAt: patient.CreatedAt,
		UpdatedAt: patient.UpdatedAt,
	}
}

// ToPatientResponseList converts slice of Patient entities to slice of PatientResponse
func ToPatientResponseList(patients []*entitiy.Patient) []*PatientResponse {
	if patients == nil {
		return nil
	}

	responses := make([]*PatientResponse, len(patients))
	for i, patient := range patients {
		responses[i] = ToPatientResponse(patient)
	}

	return responses
}

// UpdateEntity updates existing Patient entity with PatientRequest data
func (req *PatientRequest) UpdateEntity(patient *entitiy.Patient) {
	patient.FirstName = req.FirstName
	patient.LastName = req.LastName
	patient.Birthdate = req.Birthdate
	patient.Sex = req.Sex
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Email != "" {
		patient.Email = req.Email
	}
}
