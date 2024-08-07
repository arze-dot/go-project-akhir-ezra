package main

import (
	"database/sql"
	"fmt"
	"go-project-akhir-ezra/controllers"
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/middleware"
	"go-project-akhir-ezra/seeder"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	DB  *sql.DB
	err error
)

// @title Budget Tracking API
// @version 1.0
// @description This is a sample API for managing budgets, categories, and transactions.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)
	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	psqlSeeder := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlSeeder)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer DB.Close()

	// Seed the database
	seeder.SeedUsers(DB)
	fmt.Println("Successfully connected!")

	router := gin.Default()

	router.POST("/login", controllers.Login)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/transaction", controllers.GetAllTransactions)
		authorized.POST("/transaction", controllers.InsertTransaction)
		authorized.PUT("/transaction/:id", controllers.UpdateTransaction)
		authorized.DELETE("/transaction/:id", controllers.DeleteTransaction)

		authorized.GET("/category", controllers.GetAllCategories)
		authorized.POST("/category", controllers.InsertCategory)
		authorized.PUT("/category/:id", controllers.UpdateCategory)
		authorized.DELETE("/category/:id", controllers.DeleteCategory)

		authorized.GET("/user", controllers.GetAllUsers)
		authorized.POST("/user", controllers.InsertUser)
		authorized.PUT("/user/:id", controllers.UpdateUser)
		authorized.DELETE("/user/:id", controllers.DeleteUser)

		authorized.GET("/budget", controllers.GetAllBudgets)
		authorized.POST("/budget", controllers.InsertBudget)
		authorized.PUT("/budget/:id", controllers.UpdateBudget)
		authorized.DELETE("/budget/:id", controllers.DeleteBudget)

		router.GET("/budget-status", controllers.GetBudgetStatus)

	}

	router.Run(":8080")
}
