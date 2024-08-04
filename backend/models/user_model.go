package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100);uniqueIndex;not null" json:"full_name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID" json:"tasks"`
}

