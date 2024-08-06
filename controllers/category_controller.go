package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id
	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}
