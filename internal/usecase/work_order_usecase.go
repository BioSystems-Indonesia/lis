package usecase

import (
	"context"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
)

type WorkOrderUsecase interface {
	Create(ctx context.Context, req *dto.WorkOrderRequest) (*dto.WorkOrderResponse, error)
	GetByNoOrder(ctx context.Context, noOrder string) (*dto.WorkOrderResponse, error)
	Update(ctx context.Context, noOrder string, req *dto.WorkOrderRequest) (*dto.WorkOrderResponse, error)
	Delete(ctx context.Context, noOrder string) error
	GetAll(ctx context.Context) ([]*dto.WorkOrderResponse, error)
	GetByDoctor(ctx context.Context, doctor string) ([]*dto.WorkOrderResponse, error)
	GetByAnalyst(ctx context.Context, analyst string) ([]*dto.WorkOrderResponse, error)
}
