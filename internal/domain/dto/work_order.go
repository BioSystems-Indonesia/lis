package dto

type WorkOrderRequest struct {
	NoOrder  string         `json:"no_order"`
	TestCode []string       `json:"test_code"`
	Patient  PatientRequest `json:"patient"`
	Analyst  string         `json:"analyst"`
	Doctor   string         `json:"doctor"`
}
