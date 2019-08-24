package models

type Attribute struct {
	AttributeId int    `gorm:"primary_key" json:"attribute_id"`
	Name        string `gorm:"not null" json:"name"`
}

// set table name
func (Attribute) TableName() string {
	return "attribute"
}
