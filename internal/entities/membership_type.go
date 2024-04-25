package entities

import "github.com/google/uuid"

type MembershipType struct {
	ID           uuid.UUID
	Type         string 
	Description  string    
	Price        string    
	DaysDuration int
}
