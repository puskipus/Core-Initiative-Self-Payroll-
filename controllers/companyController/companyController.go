package companycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/puskipus/self-payroll/pkg/models"
)

func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&company)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": company})
}
