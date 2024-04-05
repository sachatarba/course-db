package models

import (
	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	StartDate      string `gorm:"type:DATE"`
	EndDate        string `gorm:"type:DATE"`
	MembershipType string `gorm:"type:TEXT"`
	User           User
	Gym            Gym
}
