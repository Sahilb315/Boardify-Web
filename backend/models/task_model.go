package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Description string     `gorm:"type:text" json:"description,omitempty"`
	Status      string     `gorm:"type:varchar(50);not null" json:"status"`
	Priority    string     `gorm:"type:varchar(50)" json:"priority,omitempty"`
	Deadline    *time.Time `gorm:"type:timestamp" json:"deadline,omitempty"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`
}
