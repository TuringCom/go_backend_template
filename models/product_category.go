package models

type ProductCategory struct {
	ProductId  int `gorm:"not null" json:"product_id"`
	CategoryId int `gorm:"not null" json:"category_id"`
}

// set table name
func (ProductCategory) TableName() string {
	return "product_category"
}
