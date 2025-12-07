package wire

import (
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/repository"

	"gorm.io/gorm"
)

type Repository struct {
	ProductType repository.ProductTypeRepository
	Product     repository.ProductRepository
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		ProductType: repository.NewProductTypeRepository(db),
		Product:     repository.NewProductRepository(db),
	}
}
