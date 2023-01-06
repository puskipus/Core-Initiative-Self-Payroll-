package companycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func GetDetail(c *gin.Context) {
	var company []models.Company

	if err := models.DB.Find(&company).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": company})
}

type balance struct {
	Value int `json:"balance"`
}

func TopupBalance(c *gin.Context) {
	var company models.Company
	var b balance

	balance := models.DB.First(&company).Error
	err := c.ShouldBindJSON(&b)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "jumlah balance baru harus diisi"})
		return
	}
	// fmt.Println(b.Value)

	// data := binary.BigEndian.Uint64(byteToInt)

	// newBalance, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// byteToInt, _ := strconv.Atoi(string(newBalance))
	updateBalance := company.Balance + b.Value
	// balance, _ := ioutil.ReadAll(c.Request.Body)
	// newBalance := lastBalance + balance
	// log.Printf("%+v", company)
	// log.Printf("%+v", lastBalance)
	// log.Printf("%+v", byteToInt)

	if balance != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": balance.Error()})
		return
	}

	if err := models.DB.Model(&company).Update("balance", updateBalance).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": company})

}
