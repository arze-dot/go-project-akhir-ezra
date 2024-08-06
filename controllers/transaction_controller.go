package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func InsertTransaction(c *gin.Context) {
	var transaction structs.Transaction

	err := c.BindJSON(&transaction)
	if err != nil {
		panic(err)
	}

	err = repository.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&transaction)
	if err != nil {
		panic(err)
	}

	transaction.ID = id

	err = repository.UpdateTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, _ := strconv.Atoi(c.Param("id"))

	transaction.ID = id
	err := repository.DeleteTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, transaction)
}
