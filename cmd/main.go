package main

import (
	"github.com/gin-gonic/gin"
	"github.com/puskipus/self-payroll/controllers/employeeController"
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

	router.POST("/employee", employeeController.NewEmployee)
	router.GET("/employee", employeeController.FetchEmployee)

	router.Run()

}
