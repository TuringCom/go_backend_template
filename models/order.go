package models

import "time"

// Order table model
type Order struct {
	OrderID     int     `gorm:"primary_key" json:"order_id"`
	TotalAmount float32 `gorm:"type:decimal(10,2)" json:"total_amount"`
	CreatedOn   time.Time
	ShippedOn   time.Time
	Status      int    `gorm:"not null" json:"status"`
	Comments    string `gorm:"type:varchar(255)" json:"comments"`
	CustomerID  int    `gorm:"not null" json:"customer_id"`
	AuthCode    string `gorm:"type:varchar(50)" json:"auth_code"`
	Reference   string `gorm:"type:varchar(50)" json:"reference"`
	ShippingID  int    `gorm:"not null" json:"shipping_id"`
	TaxID       int    `gorm:"not null" json:"tax_id"`
}
