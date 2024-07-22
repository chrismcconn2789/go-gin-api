package handlers

import (
	"net/http"

	"github.com/chrismconn2789/go-gin-api/db"
	"github.com/gin-gonic/gin"
)

func FindTransactions(c *gin.Context) {
	var transactions []db.Transaction
	db.DB.Find(&transactions)
	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func FindTransactionById(c *gin.Context) {
	var transaction db.Transaction
	if err := db.DB.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactionById": transaction})
}

func CreateTransaction(c *gin.Context) {
	var input db.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction := db.Transaction{Name: input.Name, Amount: input.Amount}
	db.DB.Create(&transaction)
	c.JSON(http.StatusCreated, gin.H{"createdTransaction": transaction})
}

func DeleteTransaction(c *gin.Context) {
	var transaction db.Transaction
	if err := db.DB.Where("id = ?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.DB.Delete(&transaction)
	c.JSON(http.StatusOK, gin.H{"success": "Record deleted"})
}
