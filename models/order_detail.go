package models

// OrderDetail table model
type OrderDetail struct {
	ItemID      int     `gorm:"primary_key" json:"item_id"`
	OrderID     int     `gorm:"not null" json:"order_id"`
	ProductID   int     `gorm:"not null" json:"product_id"`
	Attributes  string  `gorm:"type:varchar(1000)" json:"attributes"`
	ProductName string  `gorm:"type:varchar(100)" json:"product_name"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	UnitCost    float32 `gorm:"type:decimal(10,2)" json:"unit_cost"`
}

// TableName set table name
func (OrderDetail) TableName() string {
	return "customer"
}
