package models

import "time"

// ShoppingCart table model
type ShoppingCart struct {
	ItemID     int    `gorm:"primary_key" json:"item_id"`
	CartID     string `gorm:"type:string(32)" json:"cart_id"`
	ProductID  int    `gorm:"not null" json:"product_id"`
	Attributes string `gorm:"type:varchar(1000)" json:"attributes"`
	Quantity   int    `gorm:"not null" json:"quantity"`
	BuyNow     bool   `gorm:"type:tinyint(1);not null" json:"buy_now"`
	AddedOn    time.Time
}

// TableName set table name
func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
