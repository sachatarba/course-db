package irepositories

import (
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entities"
)

type IGymRepository interface {
	RegisterNewGym(gym entities.Gym) error
	ChangeGym(gym entities.Gym) error
	DeleteGym(gymID uuid.UUID) error
	GetGymByID(gymID uuid.UUID) (error, entities.Gym)
	ListGyms() (error, []entities.Gym)
}
