package entitiy

type WorkOrder struct {
	NoOrder   string
	PatientID string
	TestCode  []string
	Analyst   string
	Doctor    string
}
