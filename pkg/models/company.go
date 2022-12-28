package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Balance int    `gorm:"type:int" json:"balance"`
	Address string `gorm:"type:varchar(255)" json:"address"`
}
