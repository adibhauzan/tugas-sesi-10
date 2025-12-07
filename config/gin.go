package config

import (
	"net/http"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = append(configCors.AllowHeaders, "Content-Type", "Content-Length", "Accept-Encoding", "X-PARTNER-ID", "X-SIGNATUREa", "X-TIMESTAMP, X-EXTERNAL-ID", "CHANNEL-ID", "X-XSRF-TOKEN", "X-CSRF-Token", "Authorization", "X-M2M-Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers", "Access-Control-Allow-Credentials", "Origin", "Accept", "X-Requested-With", "access-control-allow-origin", "access-control-allow-methods", "access-control-allow-headers")
	configCors.AllowMethods = append(configCors.AllowMethods, "Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	configCors.AllowCredentials = true

	switch Mode {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "staging":
		gin.SetMode(gin.TestMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 30
	router.ForwardedByClientIP = true
	router.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	router.Use(cors.New(configCors))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, common.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Route Not Found",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.WebResponse{
			Code:    http.StatusOK,
			Status:  "Ok",
			Message: "Wellcome to onboarding backend provider",
		})
	})

	return router

}
