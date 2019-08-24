package models

import "time"

type Order struct {
	OrderId     int     `gorm:"primary_key" json:"order_id"`
	TotalAmount float32 `gorm:"type:decimal(10,2)" json:"total_amount"`
	CreatedOn   time.Time
	ShippedOn   time.Time
	Status      int    `gorm:"not null" json:"status"`
	Comments    string `gorm:"type:varchar(255)" json:"comments"`
	CustomerId  int    `gorm:"not null" json:"customer_id"`
	AuthCode    string `gorm:"type:varchar(50)" json:"auth_code"`
	Reference   string `gorm:"type:varchar(50)" json:"reference"`
	ShippingId  int    `gorm:"not null" json:"shipping_id"`
	TaxId       int    `gorm:"not null" json:"tax_id"`
}
