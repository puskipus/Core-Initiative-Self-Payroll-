package main

import (
	"github.com/gin-gonic/gin"
	companycontroller "github.com/puskipus/self-payroll/controllers/companyController"
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
	router.DELETE("/employee/:id", employeeController.DeleteEmployee)
	router.GET("/employee/:id", employeeController.DetailEmployee)

	router.POST("/company", companycontroller.CreateCompany)
	router.POST("/company/topup", companycontroller.TopupBalance)
	router.GET("/company", companycontroller.GetDetail)

	router.Run()

}
