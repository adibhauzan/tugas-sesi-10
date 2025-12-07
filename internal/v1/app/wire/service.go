package wire

import (
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/service"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	ProductType service.ProductTypeService
	Product     service.ProductService
}

func InitService(repo *Repository, db *gorm.DB, logger *logrus.Logger) *Service {
	return &Service{
		ProductType: service.NewProductTypeService(repo.ProductType, db, logger),
		Product:     service.NewProductService(repo.Product, repo.ProductType, db, logger),
	}
}
