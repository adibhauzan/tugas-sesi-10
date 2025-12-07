package entity

import "time"

type ProductType struct {
	ID string `gorm:"column:id"`
	// Code      int    `gorm:"column:code"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ProductType) Tablename() string {
	return "product_types"
}
