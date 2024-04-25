package entities

import "github.com/google/uuid"

type Trainer struct {
	ID            uuid.UUID
	Fullname      string 
	Email         string    
	Phone         string    
	Qualification string    
	UnitPrice     float64   
	Gyms          []*Gym    
	Trainings     []Training
}

