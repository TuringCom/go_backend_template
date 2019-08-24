package models

type AttributeValue struct {
	AttributeValueId int    `gorm:"primary_key" json:"attribute_value_id"`
	AttributeId      int    `gorm:"not null" json:"attribute_id"`
	Value            string `gorm:"type:varchar(100);not null" json:"value"`
}

// set table name
func (AttributeValue) TableName() string {
	return "attribute_value"
}
