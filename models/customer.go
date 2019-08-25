package models

// Customer table model
type Customer struct {
	CustomerID       int    `gorm:"primary_key" json:"customer_id"`
	Name             string `gorm:"type:varchar(50);not null" json:"name"`
	Email            string `gorm:"type:varchar(100);not null" json:"email"`
	Password         string `gorm:"type:varchar(100);not null" json:"password"`
	CreditCard       string
	Address1         string `gorm:"type:varchar(100)" json:"address_1"`
	Address2         string `gorm:"type:varchar(100)" json:"address_2"`
	City             string `gorm:"type:varchar(100)" json:"city"`
	Region           string `gorm:"type:varchar(100)" json:"region"`
	PostalCode       string `gorm:"type:varchar(100)" json:"postal_code"`
	Country          string `gorm:"type:varchar(100)" json:"country"`
	ShippingRegionID *int   `gorm:"default:1;not null" json:"shipping_region_id"`
	DayPhone         string `gorm:"type:varchar(100)" json:"day_phone"`
	EvePhone         string `gorm:"type:varchar(100)" json:"eve_phone"`
	MobPhone         string `gorm:"type:varchar(100)" json:"mob_phone"`
}

// TableName set table name
func (Customer) TableName() string {
	return "customer"
}
