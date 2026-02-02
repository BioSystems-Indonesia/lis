package repository

import (
	"context"
	"database/sql"

	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

type PatientRepository interface {
	Create(ctx context.Context, tx *sql.Tx, patient *entitiy.Patient) error
	GetByID(ctx context.Context, tx *sql.Tx, id string) (*entitiy.Patient, error)
	Update(ctx context.Context, tx *sql.Tx, patient *entitiy.Patient) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	GetAll(ctx context.Context, tx *sql.Tx) ([]*entitiy.Patient, error)
	Search(ctx context.Context, tx *sql.Tx, query string) ([]*entitiy.Patient, error)
}
