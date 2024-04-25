package entities

import "github.com/google/uuid"

type ClientMembership struct {
	ID               uuid.UUID
	StartDate        string    
	EndDate          string    
	MembershipType   MembershipType
}
