package models

// Shipping table model
type Shipping struct {
	ShippingID       int     `gorm:"primary_key" json:"shipping_id"`
	ShippingType     string  `gorm:"type:varchar(100)" json:"shipping_type"`
	ShippingCost     float32 `gorm:"type:decimal(10,2)" json:"shipping_cost"`
	ShippingRegionID int     `gorm:"not null" json:"shipping_region_id"`
}

// TableName set table name
func (Shipping) TableName() string {
	return "shipping_region"
}
