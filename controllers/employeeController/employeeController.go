package employeeController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/puskipus/self-payroll/pkg/models"
)

func NewEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&employee)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": employee})

}

func FetchEmployee(c *gin.Context) {
	var employee []models.Employee
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	if err := models.DB.Not("id = ?", skip).Limit(limit).Find(&employee).Model(&employee).Association("Position").Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"employee": employee})
}

func DeleteEmployee(c *gin.Context) {
	var employee []models.Employee

	id := c.Param("id")

	if err := models.DB.Delete(&employee, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak dapat dihapus"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": employee})
}

func DetailEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	if err := models.DB.First(&employee, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": employee})

}
