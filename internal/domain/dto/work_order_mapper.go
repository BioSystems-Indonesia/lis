package dto

import "github.com/BioSystems-Indonesia/lis/internal/domain/entitiy"

type WorkOrderResponse struct {
	NoOrder  string           `json:"no_order"`
	Patient  *PatientResponse `json:"patient,omitempty"`
	TestCode []string         `json:"test_code"`
	Analyst  string           `json:"analyst"`
	Doctor   string           `json:"doctor"`
}

// ToEntity converts WorkOrderRequest to WorkOrder entity
func (req *WorkOrderRequest) ToEntity(patientID string) *entitiy.WorkOrder {
	return &entitiy.WorkOrder{
		NoOrder:   req.NoOrder,
		PatientID: patientID,
		TestCode:  req.TestCode,
		Analyst:   req.Analyst,
		Doctor:    req.Doctor,
	}
}

// ToWorkOrderResponse converts WorkOrder entity to WorkOrderResponse
func ToWorkOrderResponse(workOrder *entitiy.WorkOrder, patient *entitiy.Patient) *WorkOrderResponse {
	if workOrder == nil {
		return nil
	}

	return &WorkOrderResponse{
		NoOrder:  workOrder.NoOrder,
		Patient:  ToPatientResponse(patient),
		TestCode: workOrder.TestCode,
		Analyst:  workOrder.Analyst,
		Doctor:   workOrder.Doctor,
	}
}

// ToWorkOrderResponseList converts slice of WorkOrder entities to slice of WorkOrderResponse
func ToWorkOrderResponseList(workOrders []*entitiy.WorkOrder, patients map[string]*entitiy.Patient) []*WorkOrderResponse {
	if workOrders == nil {
		return nil
	}

	responses := make([]*WorkOrderResponse, len(workOrders))
	for i, workOrder := range workOrders {
		var patient *entitiy.Patient
		if patients != nil {
			patient = patients[workOrder.PatientID]
		}
		responses[i] = ToWorkOrderResponse(workOrder, patient)
	}

	return responses
}

// UpdateEntity updates existing WorkOrder entity with WorkOrderRequest data
func (req *WorkOrderRequest) UpdateEntity(workOrder *entitiy.WorkOrder) {
	workOrder.TestCode = req.TestCode
	workOrder.Analyst = req.Analyst
	workOrder.Doctor = req.Doctor
}
