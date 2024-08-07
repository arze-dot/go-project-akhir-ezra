package controllers

import (
	"go-project-akhir-ezra/database"
	"go-project-akhir-ezra/repository"
	"go-project-akhir-ezra/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Category
// @Failure 500 {object} gin.H
// @Router /categories [get]
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

// InsertCategory godoc
// @Summary Insert a new category
// @Description Create a new category entry
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body structs.Category true "Category data"
// @Success 200 {object} structs.Category
// @Failure 500 {object} gin.H
// @Router /categories [post]
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

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Update the category entry with the specified ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body structs.Category true "Updated category data"
// @Success 200 {object} structs.Category
// @Failure 500 {object} gin.H
// @Router /categories/{id} [put]
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

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete the category entry with the specified ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} structs.Category
// @Failure 500 {object} gin.H
// @Router /categories/{id} [delete]
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
