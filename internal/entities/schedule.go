package entities

import "github.com/google/uuid"

type Schedule struct {
	ID            uuid.UUID 
	DayOfTheeWeek string    
	StartTime     string    
	EndTime       string    
	Training      Training
}
