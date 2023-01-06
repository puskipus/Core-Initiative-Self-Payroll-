package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Amount int    `gorm:"type:int" json:"amount"`
	Note   string `gorm:"type:varchar(255)" json:"note"`
	Type   string `gorm:"type:varchar(25)" json:"type"`
}
