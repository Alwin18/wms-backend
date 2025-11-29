package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Name      string         `gorm:"not null" json:"name"`
	Phone     string         `gorm:"" json:"phone"`
	Status    string         `gorm:"default:active" json:"status"` // active, inactive, suspended
	RoleID    uint           `gorm:"not null" json:"role_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Role       Role        `gorm:"foreignKey:RoleID" json:"roles,omitempty"`
	Warehouses []Warehouse `gorm:"many2many:user_warehouses;" json:"warehouses,omitempty"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}
