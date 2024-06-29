package controllers

import (
	"fmt"
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateProductResponse struct {
	Message string                 `json:"message"`
	Data    models.ProductResponse `json:"data"`
}

type DeleteProductResponse struct {
	Message string                 `json:"message"`
	Data    models.ProductResponse `json:"data"`
}

type ShowProductResponse struct {
	Message string                 `json:"message"`
	Data    models.ProductResponse `json:"data"`
}

type ListProductsResponse struct {
	Message string                   `json:"message"`
	Data    []models.ProductResponse `json:"data"`
}

type UpdateProductResponse struct {
	Message string                 `json:"message"`
	Data    models.ProductResponse `json:"data"`
}
