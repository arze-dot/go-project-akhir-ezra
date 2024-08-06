package repository

import (
	"database/sql"
	"go-project-akhir-ezra/structs"
)

func GetAllBudgets(db *sql.DB) (result []structs.Budget, err error) {
	sql := "SELECT * FROM budget"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var budget structs.Budget
		err = rows.Scan(&budget.ID, &budget.CategoryID, &budget.Price, &budget.UserID, &budget.Description, &budget.CreatedAt, &budget.CreatedBy, &budget.UpdatedAt, &budget.UpdatedBy)
		if err != nil {
			return
		}

		result = append(result, budget)
	}

	return
}

func InsertBudget(db *sql.DB, budget structs.Budget) (err error) {
	sql := "INSERT INTO budget(id, category_id, price, user_id, description, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	errs := db.QueryRow(sql, budget.ID, budget.CategoryID, budget.Price, budget.UserID, budget.Description, budget.CreatedAt, budget.CreatedBy, budget.UpdatedAt, budget.UpdatedBy)

	return errs.Err()
}

func UpdateBudget(db *sql.DB, budget structs.Budget) (err error) {
	sql := "UPDATE budget SET category_id = $1, price = $2, user_id = $3, description = $4, updated_at = $5, updated_by = $6 WHERE id = $7"

	errs := db.QueryRow(sql, budget.CategoryID, budget.Price, budget.UserID, budget.Description, budget.UpdatedAt, budget.UpdatedBy, budget.ID)

	return errs.Err()
}

func DeleteBudget(db *sql.DB, budget structs.Budget) (err error) {
	sql := "DELETE FROM budget WHERE id = $1"

	errs := db.QueryRow(sql, budget.ID)
	return errs.Err()
}
