package models

import "time"

type Position struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)" json:"name" binding:"required"`
	Salary    int    `gorm:"type:int" json:"salary" binding:"required"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
