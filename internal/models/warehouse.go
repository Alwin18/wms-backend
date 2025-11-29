package models

import (
	"time"

	"gorm.io/gorm"
)

// Warehouse represents a warehouse/distribution center
type Warehouse struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Code       string         `gorm:"uniqueIndex;not null" json:"code"`
	Name       string         `gorm:"not null" json:"name"`
	Type       string         `json:"type"` // DC, warehouse, store
	Address    string         `json:"address"`
	City       string         `json:"city"`
	Province   string         `json:"province"`
	PostalCode string         `json:"postal_code"`
	Country    string         `gorm:"default:Indonesia" json:"country"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email"`
	Status     string         `gorm:"default:active" json:"status"` // active, inactive
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Warehouse model
func (Warehouse) TableName() string {
	return "warehouses"
}
