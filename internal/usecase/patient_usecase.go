package usecase

import (
	"context"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
)

type PatientUsecase interface {
	Create(ctx context.Context, req *dto.PatientRequest) (*dto.PatientResponse, error)
	GetByID(ctx context.Context, id string) (*dto.PatientResponse, error)
	Update(ctx context.Context, id string, req *dto.PatientRequest) (*dto.PatientResponse, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*dto.PatientResponse, error)
	Search(ctx context.Context, query string) ([]*dto.PatientResponse, error)
}
