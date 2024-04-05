package models

import (
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Name        string `gorm:"type:TEXT"`
	Description string `gorm:"type:TEXT"`
	Count       int    `gorm:"type:INT;check:check_valid_count, count >= 0"`
	Gym         Gym
}
