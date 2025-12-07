package wire

import (
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/handler"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	ProductType *handler.ProductTypeHandler
	Product     *handler.ProductHandler
}

func InitHandler(svc *Service, logger *logrus.Logger) *Handler {
	return &Handler{
		ProductType: handler.NewProductTypeHandler(svc.ProductType),
		Product:     handler.NewProductHandler(svc.Product),
	}
}
