package model

import (
	"gorm.io/gorm"
)

type User struct {		
	Username string `gorm:"unique;column=username;size=256"`
	Email string `gorm:"unique;column=email;size=256"`
	PhoneNumber string `gorm:"unique;column=phone_number;size=256"`
	Password string `gorm:"column=password;size=512"`
	IsActivated bool `gorm:"column=is_activated;default=true"`
	gorm.Model
}
