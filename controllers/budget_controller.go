package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllBudgets godoc
// @Summary Get all budgets
// @Description Retrieve a list of all budgets
// @Tags budgets
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Budget
// @Failure 500 {object} gin.H
// @Router /budgets [get]
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

// InsertBudget godoc
// @Summary Insert a new budget
// @Description Create a new budget entry
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param budget body structs.Budget true "Budget data"
// @Success 200 {object} structs.Budget
// @Failure 500 {object} gin.H
// @Router /budgets [post]
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

// UpdateBudget godoc
// @Summary Update an existing budget
// @Description Update the budget entry with the specified ID
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param id path int true "Budget ID"
// @Param budget body structs.Budget true "Updated budget data"
// @Success 200 {object} structs.Budget
// @Failure 500 {object} gin.H
// @Router /budgets/{id} [put]
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

// DeleteBudget godoc
// @Summary Delete a budget
// @Description Delete the budget entry with the specified ID
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param id path int true "Budget ID"
// @Success 200 {object} structs.Budget
// @Failure 500 {object} gin.H
// @Router /budgets/{id} [delete]
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
