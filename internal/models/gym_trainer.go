package models

import (
	"gorm.io/gorm"
)

type GymTrainer struct {
	gorm.Model
	Gym     Gym
	Trainer Trainer
}
