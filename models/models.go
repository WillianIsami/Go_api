package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       string
	Stock       int
	CategoryID  uint
	Category    Category
}

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"` // "customer" or "admin"
}

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	User       User        `json:"user" gorm:"foreignKey:UserID"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"` // pending, paid, shipped, completed, canceled
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	Order     Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductId uint    `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (c Category) IsEmpty() bool {
	return c.Name == "" && len(c.Products) == 0
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deteledAt,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uint      `json:"category_id"`
}
