package models

import "time"

type ShoppingCart struct {
	ItemId     int    `gorm:"primary_key" json:"item_id"`
	CartId     string `gorm:"type:string(32)" json:"cart_id"`
	ProductId  int    `gorm:"not null" json:"product_id"`
	Attributes string `gorm:"type:varchar(1000)" json:"attributes"`
	Quantity   int    `gorm:"not null" json:"quantity"`
	BuyNow     bool   `gorm:"type:tinyint(1);not null" json:"buy_now"`
	AddedOn    time.Time
}

// set table name
func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
