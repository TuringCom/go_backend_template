package models

// Category table model
type Category struct {
	CategoryID   int    `gorm:"primary_key" json:"category_id"`
	DepartmentID int    `gorm:"not null" json:"department_id"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Description  string `gorm:"type:varchar(1000);not null" json:"description"`
}

// TableName set table name
func (Category) TableName() string {
	return "category"
}
