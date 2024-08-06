package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBudgets(c *gin.Context) {
	var (
		result gin.H
	)

	budgets, err := repository.GetAllBudgets(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": budgets,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBudget(c *gin.Context) {
	var budget structs.Budget

	err := c.BindJSON(&budget)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBudget(database.DbConnection, budget)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, budget)
}

func UpdateBudget(c *gin.Context) {
	var budget structs.Budget
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&budget)
	if err != nil {
		panic(err)
	}

	budget.ID = id

	err = repository.UpdateBudget(database.DbConnection, budget)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, budget)
}

func DeleteBudget(c *gin.Context) {
	var budget structs.Budget
	id, _ := strconv.Atoi(c.Param("id"))

	budget.ID = id
	err := repository.DeleteBudget(database.DbConnection, budget)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, budget)
}
