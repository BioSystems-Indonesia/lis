package entitiy

import "time"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Patient struct {
	ID        string
	FirstName string
	LastName  string
	Birthdate time.Time
	Sex       Gender
	Address   string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
