package models

type Category struct {
	CategoryId   int    `gorm:"primary_key" json:"category_id"`
	DepartmentId int    `gorm:"not null" json:"department_id"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Description  string `gorm:"type:varchar(1000);not null" json:"description"`
}

// set table name
func (Category) TableName() string {
	return "category"
}
