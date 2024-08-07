package repository

import (
	"database/sql"
	"go-project-akhir-ezra/structs"
)

func GetAllUsers(db *sql.DB) (result []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user structs.User
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
		if err != nil {
			return
		}

		result = append(result, user)
	}

	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users(id, username, password, role, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	errs := db.QueryRow(sql, user.ID, user.Username, user.Password, user.Role, user.CreatedAt, user.CreatedBy, user.UpdatedAt, user.UpdatedBy)

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET username = $1, password = $2, role = $3, updated_at = $4, updated_by = $5 WHERE id = $6"

	errs := db.QueryRow(sql, user.Username, user.Password, user.Role, user.UpdatedAt, user.UpdatedBy, user.ID)

	return errs.Err()
}

func DeleteUser(db *sql.DB, id int) (err error) {
	query := "DELETE FROM users WHERE id = ?"

	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var user structs.User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
