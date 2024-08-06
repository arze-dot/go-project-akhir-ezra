package structs

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}

type Transaction struct {
	ID          int       `json:"id"`
	Income      int       `json:"income"`
	Expense     int       `json:"expense"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"category,omitempty"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}

type Budget struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"category,omitempty"`
	Price       int       `json:"price"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user,omitempty"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}
