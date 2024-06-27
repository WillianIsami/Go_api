package controllers

import (
	"fmt"
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllCategory(c *gin.Context) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing categories")
		return
	}
	sendSuccess(c, "list-categories", categories)
}

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
	sendSuccess(c, "create_category", category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	category := models.Category{}
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&category)
	c.JSON(http.StatusOK, category)
}

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
