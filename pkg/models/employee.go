package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID         int      `gorm:"primaryKey"`
	SecretID   string   `gorm:"type:varchar(25)" json:"secret_id"`
	Name       string   `gorm:"type:varchar(25)" json:"name"`
	Email      string   `gorm:"type:varchar(50)" json:"email"`
	Phone      string   `gorm:"type:varchar(25)" json:"phone"`
	Address    string   `gorm:"type:varchar(50)" json:"address"`
	PositionID int      `gorm:"type:int" json:"position_id"`
	Position   Position `gorm:"references:ID"`
	CreatedAt  time.Time
	UpdateAt   time.Time
}
