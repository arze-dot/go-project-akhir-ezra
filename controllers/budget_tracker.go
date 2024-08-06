package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBudgetStatus(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
		return
	}

	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	result, err := repository.GetBudgetStatus(database.DbConnection, categoryID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category_name":    result.CategoryName,
		"budget_price":     result.BudgetPrice,
		"totalExpense":     result.TotalExpense,
		"budget_remaining": result.BudgetRemaining,
		"onTrack":          result.OnTrack,
	})
}
