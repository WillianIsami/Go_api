package controllers

import (
	"fmt"
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing products")
		return
	}
	sendSuccess(c, "list-products", products)
}

// @BasePath /api/v1

// @Summary Show product
// @Description Show a job product
// @Tags Product
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
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "product not found")
		return
	}

	sendSuccess(c, "show-product", product)
}

// @BasePath /api/v1

// @Summary Create product
// @Description Create a new job product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} CreateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	request := CreateProductRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	price := decimal.NewFromFloat(request.Price)
	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       price,
		Stock:       request.Stock,
		CategoryID:  *request.CategoryID,
	}
	if err := db.Create(&product).Error; err != nil {
		logger.Errorf("error creating product: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating product on database")
		return
	}
	if err := db.Preload("Category").First(&product, product.ID).Error; err != nil {
		logger.Errorf("error loading category: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error loading category")
		return
	}

	sendSuccess(c, "create_product", product)
}

// @BasePath /api/v1

// @Summary Update product
// @Description Update a job product
// @Tags Product
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
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		logger.Errorf("error loading category: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error loading category")
		return
	}
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		sendError(c, http.StatusNotFound, "error product not found")
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		sendError(c, http.StatusBadRequest, "error binding json")
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
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
		return
	}
	sendSuccess(c, "delete-product", product)
}
