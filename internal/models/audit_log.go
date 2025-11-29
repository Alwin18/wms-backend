package models

import (
	"time"

	"gorm.io/gorm"
)

// AuditLog represents an audit log entry
type AuditLog struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null;index" json:"user_id"`
	Action     string         `gorm:"not null" json:"action"` // e.g., "login", "create_order", "delete_product"
	Resource   string         `json:"resource"`               // e.g., "user", "product", "order"
	ResourceID string         `json:"resource_id"`
	IPAddress  string         `json:"ip_address"`
	UserAgent  string         `json:"user_agent"`
	Details    string         `gorm:"type:text" json:"details"` // JSON string for additional details
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName specifies the table name for AuditLog model
func (AuditLog) TableName() string {
	return "audit_logs"
}
