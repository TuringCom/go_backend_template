package models

// ProductCategory table model
type ProductCategory struct {
	ProductID  int `gorm:"not null" json:"product_id"`
	CategoryID int `gorm:"not null" json:"category_id"`
}

// TableName set table name
func (ProductCategory) TableName() string {
	return "product_category"
}
