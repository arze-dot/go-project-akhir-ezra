package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.User
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /user [get]
func GetAllUsers(c *gin.Context) {
	users, err := repository.GetAllUsers(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// InsertUser godoc
// @Summary Insert a new user
// @Description Create a new user entry
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body structs.User true "User data"
// @Success 200 {object} structs.User
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /user [post]
func InsertUser(c *gin.Context) {
	var user structs.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request",
		})
		return
	}

	if err := repository.InsertUser(database.DbConnection, user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update the user entry with the specified ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body structs.User true "Updated user data"
// @Success 200 {object} structs.User
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request",
		})
		return
	}

	user.ID = id

	if err := repository.UpdateUser(database.DbConnection, user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete the user entry with the specified ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} structs.User
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := repository.DeleteUser(database.DbConnection, id); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}
