package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission represents a permission for access control
type Permission struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Resource    string         `gorm:"not null" json:"resource"` // e.g., "inventory", "sales_order"
	Action      string         `gorm:"not null" json:"action"`   // e.g., "read", "create", "update", "delete"
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Permission model
func (Permission) TableName() string {
	return "permissions"
}
