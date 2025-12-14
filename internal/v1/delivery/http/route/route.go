package route

import (
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouteConfig struct {
	App                *gin.RouterGroup
	Logger             *logrus.Logger
	ProductTypeHandler *handler.ProductTypeHandler
	ProductHandler     *handler.ProductHandler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

}

func (c *RouteConfig) SetupAuthRoute() {
	// Product Type
	producType := c.App.Group("/product-type")
	producType.GET("", c.ProductTypeHandler.GetAll)
	producType.POST("create", c.ProductTypeHandler.Create)
	producType.GET(":id", c.ProductTypeHandler.GetByID)
	producType.PUT(":id/update", c.ProductTypeHandler.Update)
	producType.DELETE(":id/delete", c.ProductTypeHandler.Delete)

	// Product
	product := c.App.Group("/product")
	product.GET("", c.ProductHandler.GetAll)
	product.POST("create", c.ProductHandler.Create)
	product.GET(":id", c.ProductHandler.GetByID)
	product.PUT(":id/update", c.ProductHandler.Update)
	product.DELETE(":id/delete", c.ProductHandler.Delete)
}
