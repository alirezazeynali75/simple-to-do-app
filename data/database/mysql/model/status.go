package model

import (
	"gorm.io/gorm"
)

type Status struct {		
	Title string `gorm:"column=title;size=128"`
	Description string `gorm:"column=description;size=512"`
	Tasks []Tasks
	gorm.Model
}
