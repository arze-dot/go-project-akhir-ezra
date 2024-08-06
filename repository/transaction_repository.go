package repository

import (
	"database/sql"
	"go-project-akhir-ezra/structs"
)

func GetAllTransactions(db *sql.DB) (result []structs.Transaction, err error) {
	sql := "SELECT * FROM transaction"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var transaction structs.Transaction
		err = rows.Scan(&transaction.ID, &transaction.Income, &transaction.Expense, &transaction.CategoryID, &transaction.UserID, &transaction.Title, &transaction.Description, &transaction.CreatedAt, &transaction.CreatedBy, &transaction.UpdatedAt, &transaction.UpdatedBy)
		if err != nil {
			return
		}

		result = append(result, transaction)
	}

	return
}

func InsertTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := "INSERT INTO transaction(id, income, expense, category_id, user_id, title, description, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	errs := db.QueryRow(sql, transaction.ID, transaction.Income, transaction.Expense, transaction.CategoryID, transaction.UserID, transaction.Title, transaction.Description, transaction.CreatedAt, transaction.CreatedBy, transaction.UpdatedAt, transaction.UpdatedBy)

	return errs.Err()
}

func UpdateTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := "UPDATE transaction SET income = $1, expense = $2, category_id = $3, user_id = $4, title = $5, description = $6, updated_at = $7, updated_by = $8 WHERE id = $9"

	errs := db.QueryRow(sql, transaction.Income, transaction.Expense, transaction.CategoryID, transaction.UserID, transaction.Title, transaction.Description, transaction.UpdatedAt, transaction.UpdatedBy, transaction.ID)

	return errs.Err()
}

func DeleteTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := "DELETE FROM transaction WHERE id = $1"

	errs := db.QueryRow(sql, transaction.ID)
	return errs.Err()
}
