package models

// ShippingRegion table model
type ShippingRegion struct {
	ShippingRegionID int    `gorm:"primary_key" json:"shipping_region_id"`
	ShippingRegion   string `gorm:"type:varchar(100)" json:"shipping_region"`
}

// TableName set table name
func (ShippingRegion) TableName() string {
	return "shipping_region"
}
