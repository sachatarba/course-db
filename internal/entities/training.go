package entities

import "github.com/google/uuid"

type TrainingType int

const (
	Aerobic TrainingType = iota
	Anaerobic
	Flexibility 
	Strength
)

type Training struct {
	ID           uuid.UUID
	Title        string    
	Description  string    
	TrainingType TrainingType    
	TrainerID    uuid.UUID
}
