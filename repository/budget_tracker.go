package repository

import (
	"database/sql"
	"go-project-akhir-ezra/structs"
)

type BudgetStatus struct {
	CategoryName    string
	BudgetPrice     int
	TotalExpense    int
	BudgetRemaining int
	OnTrack         bool
}

func GetBudgetStatus(db *sql.DB, categoryID int, userID int) (result BudgetStatus, err error) {

	var category structs.Category
	categoryQuery := "SELECT name FROM category WHERE id = $1"
	err = db.QueryRow(categoryQuery, categoryID).Scan(&category.Name)
	if err != nil {
		return
	}
	result.CategoryName = category.Name

	var budget structs.Budget
	budgetQuery := "SELECT price FROM budget WHERE category_id = $1 AND user_id = $2"
	err = db.QueryRow(budgetQuery, categoryID, userID).Scan(&budget.Price)
	if err != nil {
		return
	}
	result.BudgetPrice = budget.Price

	var totalExpense int
	expenseQuery := "SELECT COALESCE(SUM(expense), 0) FROM transaction WHERE category_id = $1 AND user_id = $2"
	err = db.QueryRow(expenseQuery, categoryID, userID).Scan(&totalExpense)
	if err != nil {
		return
	}
	result.TotalExpense = totalExpense

	result.BudgetRemaining = budget.Price - totalExpense

	result.OnTrack = result.BudgetRemaining >= 0

	return
}
