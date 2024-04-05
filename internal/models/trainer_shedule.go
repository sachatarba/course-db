package models

import (
	"gorm.io/gorm"
)

type TrainerSchedule struct {
	gorm.Model
	DayOfTheeWeek string `gorm:"type:TEXT"`
	StartTime     string `gorm:"type:TIME"`
	EndTime       string `gorm:"type:TIME"`
	GymTrainer    GymTrainer
}
