package model

import (
	"gorm.io/gorm"
)

type User struct {		
	Name 				string	`gorm:"size:32;column:name"`
	FamilyName	string	`gorm:"size:32;column:family_name"`
	Email 			string	`gorm:"size:64;column:email"`
	UserName 		string	`gorm:"unique;size:64;column:user_name"`
	NationalId	uint		`gorm:"unique;size:256;column:national_id"`
	Password 		string	`gorm:"size:256;column:password"`
	IsActivated bool 		`gorm:"default:true;column:is_activated"`
	gorm.Model
}
