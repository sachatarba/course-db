package entities

import "github.com/google/uuid"

type Equipment struct {
	ID          uuid.UUID
	Name        string 
	Description string
	GymID       uuid.UUID
}
