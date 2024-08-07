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
	UserID      int       `json:"user_id"`
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
	Price       int       `json:"price"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
}

type BudgetStatus struct {
	CategoryName    string  `json:"category_name"`
	BudgetPrice     float64 `json:"budget_price"`
	TotalExpense    float64 `json:"totalExpense"`
	BudgetRemaining float64 `json:"budget_remaining"`
	OnTrack         bool    `json:"onTrack"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
