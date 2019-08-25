package models

// Attribute table model
type Attribute struct {
	AttributeID int    `gorm:"primary_key" json:"attribute_id"`
	Name        string `gorm:"not null" json:"name"`
}

// TableName set table name
func (Attribute) TableName() string {
	return "attribute"
}
