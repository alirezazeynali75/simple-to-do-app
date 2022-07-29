package model

import (
	"gorm.io/gorm"
)

type Tasks struct {		
	Title string `gorm:"column=title;size=128"`
	Summary string `gorm:"column=summary;size=512"`
	Description string `gorm:"column=description;size=512"`
	Estimation uint64 `gorm:"column=estimation;default=0"`
	StatusId uint `gorm:"column=status_id"`
	ReporterId uint `gorm:"column=reporter_id"`
	AssigneId uint `gorm:"column=assigne_id"`
	Reporter User `gorm:"foreignKey:reporter_id"`
	Assigne User `gorm:"foreignKey:assigne_id"`
	Status Status `gorm:"foreignKey:status_id"`
	gorm.Model
}
