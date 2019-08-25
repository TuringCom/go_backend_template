package models

// Product table model
type Product struct {
	ProductID       int     `gorm:"primary_key" json:"product_id"`
	Name            string  `gorm:"type:varchar(100);not null" json:"name"`
	Description     string  `gorm:"type:varchar(1000);not null" json:"description"`
	Price           float32 `gorm:"type:decimal(10,2)" json:"price"`
	DiscountedPrice float32 `gorm:"type:decimal(10,2)" json:"discounted_price"`
	Image           string  `gorm:"type:varchar(150)" json:"image"`
	Image2          string  `gorm:"type:varchar(150)" json:"image_2"`
	Thumbnail       string  `gorm:"type:varchar(150)" json:"thumbnail"`
	Display         int     `gorm:"type:smallint(6);not null" json:"status"`
}

// TableName set table name
func (Product) TableName() string {
	return "product"
}
