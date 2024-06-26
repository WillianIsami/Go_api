package controllers

import (
	"fmt"
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List products
// @Description List all job products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} ListProductsResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products := []models.Product{}
	if err := db.Find(&products).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing products")
		return
	}
	sendSuccess(c, "list-products", products)
}

// @BasePath /api/v1

// @Summary Show product
// @Description Show a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} ShowProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [get]
func GetProduct(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	product := models.Product{}
	if err := db.First(&product, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "product not found")
		return
	}

	sendSuccess(c, "show-product", product)
}

// @BasePath /api/v1

// @Summary Create product
// @Description Create a new job product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} CreateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	request := CreateProductRequest{}
	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryID:  *request.CategoryID,
		Category:    request.Category,
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// @BasePath /api/v1

// @Summary Update product
// @Description Update a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product Identification"
// @Param product body UpdateProductRequest true "Product data to Update"
// @Success 200 {object} UpdateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

// @BasePath /api/v1

// @Summary Delete product
// @Description Delete a new job product
// @Tags Product
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} DeleteProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	product := models.Product{}
	if err := db.First(&product, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
		return
	}
	sendSuccess(c, "delete-product", product)
}
