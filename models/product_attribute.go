package models

// ProductAttribute table model
type ProductAttribute struct {
	ProductID        int `gorm:"not null" json:"product_id"`
	AttributeValueID int `gorm:"not null" json:"attribute_value_id"`
}

// TableName set table name
func (ProductAttribute) TableName() string {
	return "product_attribute"
}
