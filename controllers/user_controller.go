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
// @Failure 500 {object} gin.H
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUsers(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertUser godoc
// @Summary Insert a new user
// @Description Create a new user entry
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body structs.User true "User data"
// @Success 200 {object} structs.User
// @Failure 500 {object} gin.H
// @Router /users [post]
func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
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
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = id

	err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		panic(err)
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
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id
	err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}
