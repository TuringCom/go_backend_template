package models

type Shipping struct {
	ShippingId       int     `gorm:"primary_key" json:"shipping_id"`
	ShippingType     string  `gorm:"type:varchar(100)" json:"shipping_type"`
	ShippingCost     float32 `gorm:"type:decimal(10,2)" json:"shipping_cost"`
	ShippingRegionId int     `gorm:"not null" json:"shipping_region_id"`
}

// set table name
func (Shipping) TableName() string {
	return "shipping_region"
}
