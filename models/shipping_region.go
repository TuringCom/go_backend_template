package models

type ShippingRegion struct {
	ShippingRegionId int    `gorm:"primary_key" json:"shipping_region_id"`
	ShippingRegion   string `gorm:"type:varchar(100)" json:"shipping_region"`
}

// set table name
func (ShippingRegion) TableName() string {
	return "shipping_region"
}
