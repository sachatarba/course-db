package irepositories

import (
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entities"
)

type IMembershipType interface {
	RegisterNewMembershipType(membershipType entities.MembershipType) error
	ChangeMembershipType(mem entities.Gym) error
	DeleteGym(gymID uuid.UUID) error
	GetGymByID(gymID uuid.UUID) (error, entities.Gym)
	ListGyms() (error, []entities.Gym)
}
