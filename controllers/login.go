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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "JWT token"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 401 {object} ErrorResponse "Invalid username or password"
// @Failure 500 {object} ErrorResponse "Database error or token generation failure"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginVals LoginRequest

	if err := c.BindJSON(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request"})
		return
	}

	user, err := repository.GetUserByUsername(database.DbConnection, loginVals.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, ErrorResponse{"Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{"Database error"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{"Invalid username or password"})
		return
	}

	tokenString, err := helpers.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}
