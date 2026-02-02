package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
	"github.com/BioSystems-Indonesia/lis/internal/repository"
	"github.com/google/uuid"
)

type patientUsecase struct {
	db          *sql.DB
	patientRepo repository.PatientRepository
}

func NewPatientUsecase(db *sql.DB, patientRepo repository.PatientRepository) PatientUsecase {
	return &patientUsecase{
		patientRepo: patientRepo,
	}
}

func (u *patientUsecase) Create(ctx context.Context, req *dto.PatientRequest) (*dto.PatientResponse, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	id := uuid.New().String()

	patient := req.ToEntity(id)

	if err := u.patientRepo.Create(ctx, tx, patient); err != nil {
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToPatientResponse(patient), nil
}

func (u *patientUsecase) GetByID(ctx context.Context, id string) (*dto.PatientResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	patient, err := u.patientRepo.GetByID(ctx, tx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	return dto.ToPatientResponse(patient), nil
}

func (u *patientUsecase) Update(ctx context.Context, id string, req *dto.PatientRequest) (*dto.PatientResponse, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	patient, err := u.patientRepo.GetByID(ctx, tx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	req.UpdateEntity(patient)

	if err := u.patientRepo.Update(ctx, tx, patient); err != nil {
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToPatientResponse(patient), nil
}

func (u *patientUsecase) Delete(ctx context.Context, id string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := u.patientRepo.Delete(ctx, tx, id); err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (u *patientUsecase) GetAll(ctx context.Context) ([]*dto.PatientResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	patients, err := u.patientRepo.GetAll(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all patients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToPatientResponseList(patients), nil
}

func (u *patientUsecase) Search(ctx context.Context, query string) ([]*dto.PatientResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	patients, err := u.patientRepo.Search(ctx, tx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search patients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToPatientResponseList(patients), nil
}
