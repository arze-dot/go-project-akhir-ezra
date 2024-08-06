package seeder

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedUsers seeds the database with initial user data
func SeedUsers(db *sql.DB) {
	// User data to seed
	users := []struct {
		ID        int
		Username  string
		Password  string
		Role      string
		CreatedAt string
		CreatedBy int
		UpdatedAt string
		UpdatedBy int
	}{
		{
			ID:        1,
			Username:  "admin",
			Password:  "password", // this will be hashed
			CreatedAt: "2024-08-03T14:55:00Z",
			Role:      "ADMIN",
			CreatedBy: 1,
			UpdatedAt: "2024-08-03T15:00:00Z",
			UpdatedBy: 1,
		},
	}

	for _, user := range users {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Could not hash password: %v", err)
		}

		// Insert user into the database
		_, err = db.Exec(`
			INSERT INTO users (id, username, password, created_at, created_by, updated_at, updated_by)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, user.ID, user.Username, hashedPassword, user.CreatedAt, user.CreatedBy, user.UpdatedAt, user.UpdatedBy)

		if err != nil {
			log.Fatalf("Could not insert user: %v", err)
		}
	}

	log.Println("Users seeded successfully.")
}
