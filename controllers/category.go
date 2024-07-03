package controllers

import (
	"fmt"
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
)

// @title GO_API
// @version 1.0
// @description Category management API.
// @BasePath /api/v1

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags Categories
// @Produce json
// @Success 200 {object} SuccessResponse{data=models.Category}
// @Failure 500 {object} ErrorResponse
// @Router /categories [get]
func GetAllCategory(c *gin.Context) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing categories")
		return
	}
	sendSuccess(c, "list-categories", categories)
}

// GetCategory godoc
// @Summary Get category by ID
// @Description Get a category by its ID
// @Tags Category
// @Produce json
// @Param id query string true "Category ID"
// @Success 200 {object} SuccessResponse{data=models.Category}
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /category [get]
func GetCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	category := models.Category{}
	if err := db.First(&category, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "category not found")
		return
	}

	sendSuccess(c, "show-category", category)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the given details
// @Tags Category
// @Accept json
// @Produce json
// @Param category body CreateCategoryRequest true "Category details"
// @Success 200 {object} SuccessResponse{data=models.Category}
// @Failure 500 {object} ErrorResponse
// @Router /category [post]
func CreateCategory(c *gin.Context) {
	request := CreateCategoryRequest{}
	c.BindJSON(&request)
	category := models.Category{
		Name: request.Name,
	}
	if err := db.Create(&category).Error; err != nil {
		logger.Errorf("error creating category: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating category on database")
		return
	}
	sendSuccess(c, "create-category", category)
}

// UpdateCategory godoc
// @Summary Update a existing category
// @Description update details of an existing category
// @Tags Category
// @Produce json
// @Param id query string true "Category ID"
// @Success 200 {object} SuccessResponse{data=models.Category}
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /category [put]
func UpdateCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	category := models.Category{}
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		sendError(c, http.StatusBadRequest, "error category not found")
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		sendError(c, http.StatusBadRequest, "error binding json")
		return
	}
	db.Save(&category)
	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Delete a category by its ID
// @Tags Category
// @Param id query string true "Category ID"
// @Success 200 {object} SuccessResponse{data=models.Category}
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /category [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	category := models.Category{}
	if err := db.First(&category, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("category with id: %s not found", id))
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting category with id: %s", id))
		return
	}
	sendSuccess(c, "delete-category", category)
}
