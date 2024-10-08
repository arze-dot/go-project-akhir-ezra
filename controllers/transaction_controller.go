package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve a list of all transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Transaction
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /transaction [get]
func GetAllTransactions(c *gin.Context) {
	var (
		result gin.H
	)

	transactions, err := repository.GetAllTransactions(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": transactions,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertTransaction godoc
// @Summary Insert a new transaction
// @Description Create a new transaction entry
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body structs.Transaction true "Transaction data"
// @Success 200 {object} structs.Transaction
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /transaction [post]
func InsertTransaction(c *gin.Context) {
	var transaction structs.Transaction

	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if (transaction.Income != 0 && transaction.Expense != 0) || (transaction.Income == 0 && transaction.Expense == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "One of income or expense must be 0"})
		return
	}

	err = repository.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction godoc
// @Summary Update an existing transaction
// @Description Update the transaction entry with the specified ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Param transaction body structs.Transaction true "Updated transaction data"
// @Success 200 {object} structs.Transaction
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /transaction/{id} [put]
func UpdateTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	transaction.ID = id

	err = repository.UpdateTransaction(database.DbConnection, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete the transaction entry with the specified ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} structs.Transaction
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /transaction/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, _ := strconv.Atoi(c.Param("id"))

	transaction.ID = id
	err := repository.DeleteTransaction(database.DbConnection, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
