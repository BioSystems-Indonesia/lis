package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

type PatientRepositoryImpl struct{}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &PatientRepositoryImpl{}
}

func (r *PatientRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, patient *entitiy.Patient) error {
	query := `
		INSERT INTO patients (id, first_name, last_name, birthdate, sex, address, phone, email)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(ctx, query,
		patient.ID,
		patient.FirstName,
		patient.LastName,
		patient.Birthdate.Format("2006-01-02"),
		patient.Sex,
		patient.Address,
		patient.Phone,
		patient.Email,
	)

	if err != nil {
		return fmt.Errorf("failed to create patient: %w", err)
	}

	return nil
}

func (r *PatientRepositoryImpl) GetByID(ctx context.Context, tx *sql.Tx, id string) (*entitiy.Patient, error) {
	query := `
		SELECT id, first_name, last_name, birthdate, sex, address, phone, email, created_at, updated_at
		FROM patients
		WHERE id = ?
	`

	patient := &entitiy.Patient{}

	err := tx.QueryRowContext(ctx, query, id).Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.Birthdate,
		&patient.Sex,
		&patient.Address,
		&patient.Phone,
		&patient.Email,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("patient not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	return patient, nil
}

func (r *PatientRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, patient *entitiy.Patient) error {
	query := `
		UPDATE patients
		SET first_name = ?, birthdate = ?, sex = ?, address = ?, phone = ?, email = ?
		WHERE id = ?
	`

	result, err := tx.ExecContext(ctx, query,
		patient.FirstName,
		patient.Birthdate.Format("2006-01-02"),
		patient.Sex,
		patient.Address,
		patient.Phone,
		patient.Email,
		patient.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("patient not found")
	}

	return nil
}

func (r *PatientRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := `DELETE FROM patients WHERE`

	result, err := tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("patient not found")
	}

	return nil
}

func (r *PatientRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]*entitiy.Patient, error) {
	query := `
		SELECT id, first_name, last_name, birthdate, sex, address, phone, email, 
		FROM patients
		ORDER BY first_name, last_name
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}
	defer rows.Close()

	var patients []*entitiy.Patient

	for rows.Next() {
		patient := &entitiy.Patient{}

		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.Birthdate,
			&patient.Address,
			&patient.Phone,
			&patient.Email,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan patient: %w", err)
		}

		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating patients: %w", err)
	}

	return patients, nil
}

func (r *PatientRepositoryImpl) Search(ctx context.Context, tx *sql.Tx, query string) ([]*entitiy.Patient, error) {
	searchQuery := `
		SELECT id, first_name, last_name, sex, address, phone, email, created_at, updated_at
		FROM patients
		WHERE first_name LIKE ? OR last_name LIKE ? OR phone LIKE ? OR email LIKE ?
		ORDER BY first_name, last_name
	`

	searchPattern := "%" + query + "%"

	rows, err := tx.QueryContext(ctx, searchQuery, searchPattern, searchPattern, searchPattern, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search patients: %w", err)
	}
	defer rows.Close()

	var patients []*entitiy.Patient

	for rows.Next() {
		patient := &entitiy.Patient{}

		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.Birthdate,
			&patient.Sex,
			&patient.Address,
			&patient.Phone,
			&patient.Email,
			&patient.CreatedAt,
			&patient.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan patient: %w", err)
		}

		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating patients: %w", err)
	}

	return patients, nil
}
