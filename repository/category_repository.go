package repository

import (
	"database/sql"
	"go-project-akhir-ezra/structs"
)

func GetAllCategories(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.CreatedBy, &category.UpdatedAt, &category.UpdatedBy)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category(id, name, description, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	errs := db.QueryRow(sql, category.ID, category.Name, category.Description, category.CreatedAt, category.CreatedBy, category.UpdatedAt, category.UpdatedBy)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, description = $2, updated_at = $3, updated_by = $4 WHERE id = $5"

	errs := db.QueryRow(sql, category.Name, category.Description, category.UpdatedAt, category.UpdatedBy, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}
