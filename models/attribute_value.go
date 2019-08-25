package models

// AttributeValue table model
type AttributeValue struct {
	AttributeValueID int    `gorm:"primary_key" json:"attribute_value_id"`
	AttributeID      int    `gorm:"not null" json:"attribute_id"`
	Value            string `gorm:"type:varchar(100);not null" json:"value"`
}

// TableName set table name
func (AttributeValue) TableName() string {
	return "attribute_value"
}
