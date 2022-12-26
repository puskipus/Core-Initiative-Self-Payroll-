package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Salary    int    `gorm:"type:int" json:"salary"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
