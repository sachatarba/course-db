package irepositories

import (
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entities"
)

type IEquipmentRepository interface {
	CreateNewEquipment(equipment entities.Equipment) error
	ChangeEquipment(equipment entities.Equipment) error
	DeleteEquipment(equipmentID uuid.UUID) error
	GetEquipmentByID(equipmentID uuid.UUID) (error, entities.Equipment)
	ListEquipmentsByGymID(gymID uuid.UUID) (error, []entities.Equipment)
}
