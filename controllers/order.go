package controllers

import (
	"net/http"

	"github.com/WillianIsami/go_api/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreateOrderInput struct {
	Total      float64                `json:"total" binding:"required"`
	Status     string                 `json:"status" binding:"required"`
	OrderItems []CreateOrderItemInput `json:"order_items" binding:"required"`
}

type CreateOrderItemInput struct {
	OrderID   uint   `json:"order_id" binding:"required"`
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Price     string `json:"price" binding:"required"`
}

func GetAllOrders(c *gin.Context) {
	orders := []models.Order{}
	if err := db.Find(&orders).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing orders")
		return
	}
	sendSuccess(c, "list-orders", orders)
}

func GetOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	order := models.Order{}
	if err := db.First(&order, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "error order not found")
		return
	}
	sendSuccess(c, "show-order", order)
}

func CreateOrder(c *gin.Context) {
	input := CreateOrderInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		sendError(c, http.StatusBadRequest, "error creating order while binding json")
		return
	}
	orderItems := []models.OrderItem{}
	for _, item := range input.OrderItems {
		price, err := decimal.NewFromString(item.Price)
		if err != nil {
			sendError(c, http.StatusInternalServerError, "invalid price format")
			return
		}
		orderItem := models.OrderItem{
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     price,
		}
		orderItems = append(orderItems, orderItem)
	}
	order := models.Order{
		Total:      input.Total,
		Status:     input.Status,
		OrderItems: orderItems,
	}
	if err := db.Create(&order).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error creating order")
		return
	}
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}
	if err := db.Save(&orderItems).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error saving order items")
		return
	}
	sendSuccess(c, "create-order", order)
}
