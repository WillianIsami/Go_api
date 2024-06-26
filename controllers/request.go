package controllers

import (
	"fmt"

	"github.com/WillianIsami/go_api/models"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateProductRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       string          `json:"price"`
	Stock       int             `json:"stock"`
	CategoryID  *uint           `json:"category_id"`
	Category    models.Category `json:"category" gorm:"foreignKey:CategoryID"`
}

func (r *CreateProductRequest) Validate() error {
	conditions := map[string]bool{
		"name":        r.Name == "",
		"description": r.Description == "",
		"price":       r.Price == "",
		"stock":       r.Stock <= 0,
		"category_id": r.CategoryID == nil,
		"category":    r.Category.IsEmpty(),
	}

	if r.Name == "" && r.Description == "" && r.Price == "" && r.Stock <= 0 && r.CategoryID == nil && r.Category.IsEmpty() {
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
		if v {
			return errParamIsRequired(k, "string")
		}
	}
	return nil
}

type UpdateProductRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       string          `json:"price"`
	Stock       int             `json:"stock"`
	CategoryID  *uint           `json:"category_id"`
	Category    models.Category `json:"category" gorm:"foreignKey:CategoryID"`
}

func (r *UpdateProductRequest) Validate() error {
	if r.Name == "" && r.Description == "" && r.Price == "" && r.Stock <= 0 && r.CategoryID == nil && r.Category.IsEmpty() {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
