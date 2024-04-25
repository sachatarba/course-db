package entities

import (
	"github.com/google/uuid"
)

type Client struct {
	ID                uuid.UUID 
	Fullname          string    
	Email             string    
	Phone             string    
	Birthdate         string
	ClientMemberships []ClientMembership
	Schedules         []Schedule
}
