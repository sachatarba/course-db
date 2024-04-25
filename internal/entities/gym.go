package entities

import "github.com/google/uuid"

type Gym struct {
	ID              uuid.UUID
	Name            string  
	Phone           string
	City            string
	Addres          string
	IsChain         bool
	Trainers        []Trainer 
	Equipments      []Equipment
	MembershipTypes []MembershipType
}

