package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123 dbname=self_payroll port=5432 sslmode=disable TimeZone=Asia/Shanghai"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Position{}, &Employee{}, &Company{})

	DB = database
}
