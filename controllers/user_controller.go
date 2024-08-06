package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
