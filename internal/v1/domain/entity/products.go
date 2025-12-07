package entity

import "time"

type Product struct {
	ID        string      `gorm:"column:id;primaryKey"`
	Code      int         `gorm:"column:code"`
	Name      string      `gorm:"column:name"`
	TypeID    string      `gorm:"column:type_id"`
	Type      ProductType `gorm:"foreignKey:TypeID;references:ID"`
	Price     float64     `gorm:"column:price"`
	Stock     int         `gorm:"column:stock"`
	CreatedAt time.Time   `gorm:"column:created_at"`
	UpdatedAt time.Time   `gorm:"column:updated_at"`
}

func (Product) Tablename() string {
	return "master_product"
}
