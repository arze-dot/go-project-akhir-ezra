package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBudgetStatus godoc
// @Summary Get budget status
// @Description Retrieve the budget status for a given user and category
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param category_id query int true "Category ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} structs.BudgetStatus "Budget status"
// @Failure 400 {object} structs.ErrorResponse "Invalid input"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Security Bearer
// @Router /budget-status [get]
func GetBudgetStatus(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: "Invalid category_id"})
		return
	}

	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: "Invalid user_id"})
		return
	}

	result, err := repository.GetBudgetStatus(database.DbConnection, categoryID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, structs.BudgetStatus{
		CategoryName:    result.CategoryName,
		BudgetPrice:     float64(result.BudgetPrice),
		TotalExpense:    float64(result.TotalExpense),
		BudgetRemaining: float64(result.BudgetRemaining),
		OnTrack:         result.OnTrack,
	})
}
