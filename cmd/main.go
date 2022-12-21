package main

import (
	"github.com/gin-gonic/gin"
	positioncontroller "github.com/puskipus/self-payroll/controllers/positionController"
	"github.com/puskipus/self-payroll/pkg/models"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.POST("/positions", positioncontroller.NewPosition)
	router.GET("/positions", positioncontroller.FetchPosition)
	router.GET("/positions/:id", positioncontroller.DetailPosition)
	router.DELETE("/positions/:id", positioncontroller.DeletePosition)

	router.Run()

}
