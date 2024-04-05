package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:TEXT"`
	Email     string `gorm:"type:TEXT;check:check_valid_email, (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}$')"`
	Phone     string `gorm:"type:TEXT;check:check_valid_international_phone, (phone ~ '^\\+[0-9]+-[0-9]+-[0-9]+$')"`
	Birthdate string `gorm:"type:DATE"`
}
