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
// @Failure 500 {object} structs.ErrorResponse
// @Security Bearer
// @Router /category [get]
func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// InsertCategory godoc
// @Summary Insert a new category
// @Description Create a new category entry
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body structs.Category true "Category data"
// @Success 200 {object} structs.Category
// @Failure 500 {object} structs.ErrorResponse
// @Security Bearer
// @Router /category [post]
func InsertCategory(c *gin.Context) {
	var category structs.Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: "Invalid input"})
		return
	}

	if err := repository.InsertCategory(database.DbConnection, category); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{Error: err.Error()})
		return
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
// @Failure 400 {object} structs.ErrorResponse
// @Failure 500 {object} structs.ErrorResponse
// @Security Bearer
// @Router /category/{id} [put]
func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{Error: "Invalid input"})
		return
	}

	category.ID = id

	if err := repository.UpdateCategory(database.DbConnection, category); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{Error: err.Error()})
		return
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
// @Failure 500 {object} structs.ErrorResponse
// @Security Bearer
// @Router /category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := repository.DeleteCategory(database.DbConnection, structs.Category{ID: id}); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
