package app

import (
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/app/wire"
	route "tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/delivery"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Engine *gin.RouterGroup
	Log    *logrus.Logger
}

func (app *App) Run() {
	repo := wire.InitRepository(app.DB)
	svc := wire.InitService(repo, app.DB, app.Log)
	h := wire.InitHandler(svc, app.Log)
	InitRoute(app.Engine, h, svc, app.Log)
}

func InitRoute(app *gin.RouterGroup, h *wire.Handler, svc *wire.Service, logger *logrus.Logger) {
	v1 := app.Group("v1")

	routeConfig := route.RouteConfig{
		App:                v1,
		ProductTypeHandler: h.ProductType,
		ProductHandler:     h.Product,
		Logger:             logger,
	}

	routeConfig.Setup()
}
