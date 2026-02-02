package repository

import (
	"context"
	"database/sql"

	"github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"
)

type WorkOrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, workOrder *entitiy.WorkOrder) error
	GetByNoOrder(ctx context.Context, tx *sql.Tx, noOrder string) (*entitiy.WorkOrder, error)
	Update(ctx context.Context, tx *sql.Tx, workOrder *entitiy.WorkOrder) error
	Delete(ctx context.Context, tx *sql.Tx, noOrder string) error
	GetAll(ctx context.Context, tx *sql.Tx) ([]*entitiy.WorkOrder, error)
	GetByDoctor(ctx context.Context, tx *sql.Tx, doctor string) ([]*entitiy.WorkOrder, error)
	GetByAnalyst(ctx context.Context, tx *sql.Tx, analyst string) ([]*entitiy.WorkOrder, error)
}
