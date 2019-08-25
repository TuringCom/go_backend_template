package models

import "time"

// Audit table model
type Audit struct {
	AuditID   int `gorm:"primary_key" json:"audit_id"`
	OrderID   int `gorm:"not null" json:"order_id"`
	CreatedOn time.Time
	Message   string `gorm:"not null" json:"message"`
	Code      int    `gorm:"not null" json:"code"`
}

// TableName set table name
func (Audit) TableName() string {
	return "audit"
}
