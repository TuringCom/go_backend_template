package models

type Department struct {
	DepartmentId int    `gorm:"primary_key" json:"department_id"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Description  string `gorm:"type:varchar(1000)" json:"description"`
}

// set table name
func (Department) TableName() string {
	return "department"
}
