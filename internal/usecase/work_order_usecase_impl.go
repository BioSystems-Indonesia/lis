package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
	"github.com/BioSystems-Indonesia/lis/internal/repository"
	"github.com/google/uuid"
)

type workOrderUsecase struct {
	db            *sql.DB
	workOrderRepo repository.WorkOrderRepository
	patientRepo   repository.PatientRepository
}

func NewWorkOrderUsecase(db *sql.DB, workOrderRepo repository.WorkOrderRepository, patientRepo repository.PatientRepository) WorkOrderUsecase {
	return &workOrderUsecase{
		db:            db,
		workOrderRepo: workOrderRepo,
	}
}

func (u *workOrderUsecase) Create(ctx context.Context, req *dto.WorkOrderRequest) (*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	patientID := uuid.New().String()
	patient := req.Patient.ToEntity(patientID)

	if err := u.patientRepo.Create(ctx, tx, patient); err != nil {
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}

	workOrder := req.ToEntity(patientID)

	if err := u.workOrderRepo.Create(ctx, tx, workOrder); err != nil {
		return nil, fmt.Errorf("failed to create work order: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponse(workOrder, patient), nil
}

func (u *workOrderUsecase) GetByNoOrder(ctx context.Context, noOrder string) (*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	workOrder, err := u.workOrderRepo.GetByNoOrder(ctx, tx, noOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to get work order: %w", err)
	}

	patient, err := u.patientRepo.GetByID(ctx, tx, workOrder.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponse(workOrder, patient), nil
}

func (u *workOrderUsecase) Update(ctx context.Context, noOrder string, req *dto.WorkOrderRequest) (*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	workOrder, err := u.workOrderRepo.GetByNoOrder(ctx, tx, noOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to get work order: %w", err)
	}

	patient, err := u.patientRepo.GetByID(ctx, tx, workOrder.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	req.Patient.UpdateEntity(patient)
	if err := u.patientRepo.Update(ctx, tx, patient); err != nil {
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}

	req.UpdateEntity(workOrder)

	if err := u.workOrderRepo.Update(ctx, tx, workOrder); err != nil {
		return nil, fmt.Errorf("failed to update work order: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponse(workOrder, patient), nil
}

func (u *workOrderUsecase) Delete(ctx context.Context, noOrder string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := u.workOrderRepo.Delete(ctx, tx, noOrder); err != nil {
		return fmt.Errorf("failed to delete work order: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (u *workOrderUsecase) GetAll(ctx context.Context) ([]*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	workOrders, err := u.workOrderRepo.GetAll(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all work orders: %w", err)
	}

	patients, err := u.getPatientsForWorkOrders(ctx, tx, workOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponseList(workOrders, patients), nil
}

func (u *workOrderUsecase) GetByDoctor(ctx context.Context, doctor string) ([]*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	workOrders, err := u.workOrderRepo.GetByDoctor(ctx, tx, doctor)
	if err != nil {
		return nil, fmt.Errorf("failed to get work orders by doctor: %w", err)
	}

	patients, err := u.getPatientsForWorkOrders(ctx, tx, workOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponseList(workOrders, patients), nil
}

func (u *workOrderUsecase) GetByAnalyst(ctx context.Context, analyst string) ([]*dto.WorkOrderResponse, error) {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	workOrders, err := u.workOrderRepo.GetByAnalyst(ctx, tx, analyst)
	if err != nil {
		return nil, fmt.Errorf("failed to get work orders by analyst: %w", err)
	}

	patients, err := u.getPatientsForWorkOrders(ctx, tx, workOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return dto.ToWorkOrderResponseList(workOrders, patients), nil
}

func (u *workOrderUsecase) getPatientsForWorkOrders(ctx context.Context, tx *sql.Tx, workOrders []*entitiy.WorkOrder) (map[string]*entitiy.Patient, error) {
	patients := make(map[string]*entitiy.Patient)

	for _, wo := range workOrders {
		patient, err := u.patientRepo.GetByID(ctx, tx, wo.PatientID)
		if err != nil {
			continue
		}
		patients[wo.PatientID] = patient
	}

	return patients, nil
}
