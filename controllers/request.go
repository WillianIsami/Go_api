package controllers

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  *uint   `json:"category_id" binding:"required"`
}

func (r *CreateProductRequest) Validate() error {
	price := decimal.NewFromFloat(r.Price)
	conditions := map[string]bool{
		"name":        r.Name == "",
		"description": r.Description == "",
		"price":       !price.GreaterThan(decimal.Zero),
		"stock":       r.Stock <= 0,
		"category_id": r.CategoryID == nil,
	}

	if r.Name == "" && r.Description == "" && !price.GreaterThan(decimal.Zero) && r.Stock <= 0 && r.CategoryID == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	for k, v := range conditions {
		if conditions["stock"] {
			return errParamIsRequired(k, "int")
		}
		if conditions["category_id"] {
			return errParamIsRequired(k, "uint")
		}
		if conditions["category"] {
			return errParamIsRequired(k, "Category")
		}
		if conditions["price"] {
			return errParamIsRequired(k, "float64")
		}
		if v {
			return errParamIsRequired(k, "string")
		}
	}
	return nil
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  *uint   `json:"category_id" binding:"required"`
}

func (r *UpdateProductRequest) Validate() error {
	price := decimal.NewFromFloat(r.Price)
	if r.Name == "" && r.Description == "" && !price.GreaterThan(decimal.Zero) && r.Stock <= 0 && r.CategoryID == nil {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
