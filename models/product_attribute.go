package models

type ProductAttribute struct {
	ProductId        int `gorm:"not null" json:"product_id"`
	AttributeValueId int `gorm:"not null" json:"attribute_value_id"`
}

// set table name
func (ProductAttribute) TableName() string {
	return "product_attribute"
}
