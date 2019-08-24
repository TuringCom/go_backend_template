package models

import "time"

type Audit struct {
	AuditId   int `gorm:"primary_key" json:"audit_id"`
	OrderId   int `gorm:"not null" json:"order_id"`
	CreatedOn time.Time
	Message   string `gorm:"not null" json:"message"`
	Code      int    `gorm:"not null" json:"code"`
}

// set table name
func (Audit) TableName() string {
	return "audit"
}
