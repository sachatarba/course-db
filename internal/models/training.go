package models

import (
	"gorm.io/gorm"
)

type Training struct {
	gorm.Model
	Title       string `gorm:"type:TEXT"`
	Description string `gorm:"type:TEXT"`
	Trainer     Trainer
}
