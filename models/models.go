package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string          `gorm:"size:255" json:"name"`
	Description string          `gorm:"size:255" json:"description"`
	Price       decimal.Decimal `gorm:"type:decimal(10,2)" json:"price"`
	Stock       int             `json:"stock"`
	CategoryID  uint            `json:"category_id"`
	Category    Category        `gorm:"foreignKey:CategoryID"`
}

type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string `gorm:"size:255;uniqueIndex" json:"name"`
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
	OrderID   uint            `json:"order_id"`
	Order     Order           `json:"order" gorm:"foreignKey:OrderID"`
	ProductId uint            `json:"product_id"`
	Product   Product         `json:"product"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
}

type ProductResponse struct {
	ID          uint            `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   time.Time       `json:"deteledAt,omitempty"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
	Stock       int             `json:"stock"`
	CategoryID  uint            `json:"category_id"`
	Category    Category        `json:"category"`
}
