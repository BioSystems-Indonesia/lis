package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

type WorkOrderRepositoryImpl struct{}

func NewWorkOrderRepository(db *sql.DB) WorkOrderRepository {
	return &WorkOrderRepositoryImpl{}
}

func (r *WorkOrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, workOrder *entitiy.WorkOrder) error {
	query := `
		INSERT INTO work_orders (no_order, patient_id, analyst, doctor)
		VALUES (?, ?, ?, ?)
	`

	_, err := tx.ExecContext(ctx, query,
		workOrder.NoOrder,
		workOrder.PatientID,
		workOrder.Analyst,
	)

	if err != nil {
		return fmt.Errorf("failed to create work order: %w", err)
	}

	if len(workOrder.TestCode) > 0 {
		testCodeQuery := `INSERT INTO work_order_test_code (no_order, test_code) VALUES (?, ?)`

		for _, testCode := range workOrder.TestCode {
			_, err = tx.ExecContext(ctx, testCodeQuery, workOrder.NoOrder, testCode)
			if err != nil {
				return fmt.Errorf("failed to insert test code: %w", err)
			}
		}
	}

	return nil
}

func (r *WorkOrderRepositoryImpl) GetByNoOrder(ctx context.Context, tx *sql.Tx, noOrder string) (*entitiy.WorkOrder, error) {
	query := `
		SELECT no_order, patient_id, analyst, doctor
		FROM work_orders
		WHERE no_order = ?
	`

	workOrder := &entitiy.WorkOrder{}

	err := tx.QueryRowContext(ctx, query, noOrder).Scan(
		&workOrder.NoOrder,
		&workOrder.PatientID,
		&workOrder.Analyst,
		&workOrder.Doctor,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("work order not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get work order: %w", err)
	}

	testCodeQuery := `
		SELECT test_code
		FROM work_order_test_codes
		ORDER BY id
	`

	rows, err := tx.QueryContext(ctx, testCodeQuery, noOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to get test codes: %w", err)
	}

	var testCodes []string
	for rows.Next() {
		var testCode string
		if err := rows.Scan(&testCode); err != nil {
			return nil, fmt.Errorf("failed to scan test code: %w", err)
		}
		testCodes = append(testCodes, testCode)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating test codes: %w", err)
	}

	workOrder.TestCode = testCodes

	return workOrder, nil
}

func (r *WorkOrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, workOrder *entitiy.WorkOrder) error {
	query := `
		UPDATE work_orders
		SET analyst = ?, doctor = ?
		WHERE no_order = ?
	`

	result, err := tx.ExecContext(ctx, query,
		workOrder.Analyst,
		workOrder.Doctor,
		workOrder.NoOrder,
	)

	if err != nil {
		return fmt.Errorf("failed to update work order: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("work order not found")
	}

	deleteTestCodesQuery := `DELETE FROM work_order_test_codes`
	_, err = tx.ExecContext(ctx, deleteTestCodesQuery, workOrder.NoOrder)
	if err != nil {
		return fmt.Errorf("failed to delete test codes: %w", err)
	}

	if len(workOrder.TestCode) > 0 {
		testCodeQuery := `INSERT INTO work_order_test_codes (no_order, test_code) VALUES (?, ?)`

		for _, testCode := range workOrder.TestCode {
			_, err = tx.ExecContext(ctx, testCodeQuery, workOrder.NoOrder, testCode)
			if err != nil {
				return fmt.Errorf("failed to insert test code: %w", err)
			}
		}
	}

	return nil
}

func (r *WorkOrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, noOrder string) error {
	deleteTestCodesQuery := `DELETE FROM work_order_test_codes WHERE no_order = ?`
	_, err := tx.ExecContext(ctx, deleteTestCodesQuery, noOrder)
	if err != nil {
		return fmt.Errorf("failed to delete test codes: %w", err)
	}

	query := `DELETE FROM work_orders WHERE no_order = ?`

	result, err := tx.ExecContext(ctx, query, noOrder)
	if err != nil {
		return fmt.Errorf("failed to delete work order: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("work order not found")
	}

	return nil
}

func (r *WorkOrderRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]*entitiy.WorkOrder, error) {
	query := `
		SELECT no_order, patient_id, analyst, doctor
		FROM work_orders
		ORDER BY no_order
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get work orders: %w", err)
	}

	var workOrders []*entitiy.WorkOrder

	for rows.Next() {
		workOrder := &entitiy.WorkOrder{}

		err := rows.Scan(
			&workOrder.NoOrder,
			&workOrder.PatientID,
			&workOrder.Analyst,
			&workOrder.Doctor,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan work order: %w", err)
		}

		workOrders = append(workOrders, workOrder)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating work orders: %w", err)
	}

	for _, workOrder := range workOrders {
		testCodes, err := r.getTestCodes(ctx, tx, workOrder.NoOrder)
		if err != nil {
			return nil, err
		}
		workOrder.TestCode = testCodes
	}

	return workOrders, nil
}

func (r *WorkOrderRepositoryImpl) GetByDoctor(ctx context.Context, tx *sql.Tx, doctor string) ([]*entitiy.WorkOrder, error) {
	query := `
		SELECT no_order, patient_id, analyst, doctor
		FROM work_orders
		WHERE doctor = ?
		ORDER BY no_order
	`

	rows, err := tx.QueryContext(ctx, query, doctor)
	if err != nil {
		return nil, fmt.Errorf("failed to get work orders by doctor: %w", err)
	}
	defer rows.Close()

	var workOrders []*entitiy.WorkOrder

	for rows.Next() {
		workOrder := &entitiy.WorkOrder{}

		err := rows.Scan(
			&workOrder.NoOrder,
			&workOrder.PatientID,
			&workOrder.Analyst,
			&workOrder.Doctor,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan work order: %w", err)
		}

		workOrders = append(workOrders, workOrder)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating work orders: %w", err)
	}

	for _, workOrder := range workOrders {
		testCodes, err := r.getTestCodes(ctx, tx, workOrder.NoOrder)
		if err != nil {
			return nil, err
		}
		workOrder.TestCode = testCodes
	}

	return workOrders, nil
}

func (r *WorkOrderRepositoryImpl) GetByAnalyst(ctx context.Context, tx *sql.Tx, analyst string) ([]*entitiy.WorkOrder, error) {
	query := `
		SELECT no_order, patient_id, analyst, doctor
		FROM work_orders
		WHERE analyst = ?
		ORDER BY no_order
	`

	rows, err := tx.QueryContext(ctx, query, analyst)
	if err != nil {
		return nil, fmt.Errorf("failed to get work orders by analyst: %w", err)
	}
	defer rows.Close()

	var workOrders []*entitiy.WorkOrder

	for rows.Next() {
		workOrder := &entitiy.WorkOrder{}

		err := rows.Scan(
			&workOrder.NoOrder,
			&workOrder.PatientID,
			&workOrder.Analyst,
			&workOrder.Doctor,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan work order: %w", err)
		}

		workOrders = append(workOrders, workOrder)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating work orders: %w", err)
	}

	for _, workOrder := range workOrders {
		testCodes, err := r.getTestCodes(ctx, tx, workOrder.NoOrder)
		if err != nil {
			return nil, err
		}
		workOrder.TestCode = testCodes
	}

	return workOrders, nil
}

func (r *WorkOrderRepositoryImpl) getTestCodes(ctx context.Context, tx *sql.Tx, noOrder string) ([]string, error) {
	query := `
		SELECT test_code
		FROM work_order_test_codes
		WHERE no_order = ?
		ORDER BY id
	`

	rows, err := tx.QueryContext(ctx, query, noOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to get test codes: %w", err)
	}
	defer rows.Close()

	var testCodes []string
	for rows.Next() {
		var testCode string
		if err := rows.Scan(&testCode); err != nil {
			return nil, fmt.Errorf("failed to scan test code: %w", err)
		}
		testCodes = append(testCodes, testCode)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating test codes: %w", err)
	}

	return testCodes, nil
}
