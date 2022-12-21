package positioncontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/puskipus/self-payroll/pkg/models"
	"gorm.io/gorm"
)

func NewPosition(c *gin.Context) {
	var position models.Position

	if err := c.ShouldBindJSON(&position); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&position)
	c.JSON(http.StatusOK, gin.H{"position": position})

}

func DetailPosition(c *gin.Context) {
	var position models.Position
	id := c.Param("id")

	if err := models.DB.First(&position, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"position": position})

}

func FetchPosition(c *gin.Context) {
	var positions []models.Position
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	if err := models.DB.Not("id = ?", skip).Limit(limit).Find(&positions).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"positions": positions})
}

func DeletePosition(c *gin.Context) {
	var position []models.Position

	id := c.Param("id")

	if err := models.DB.Delete(&position, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak dapat dihapus"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}
