package transactioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/puskipus/self-payroll/pkg/models"
)

func GetTransactions(c *gin.Context) {
	var transaction []models.Transaction

	if err := models.DB.Find(&transaction).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": transaction})

}
