package controllers

import (
	"database/sql"
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/helpers"
	"go-project-akhir-ezra/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginVals struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repository.GetUserByUsername(database.DbConnection, loginVals.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenString, err := helpers.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
