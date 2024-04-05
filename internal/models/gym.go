package models

import (
	"gorm.io/gorm"
)

type Gym struct {
	gorm.Model
	Name     string `gorm:"type:TEXT"`
	Capacity int    `gorm:"type:INT;check:check_valid_capacity, count >= 0"`
	Addres   string `gorm:"type:TEXT"`
}
