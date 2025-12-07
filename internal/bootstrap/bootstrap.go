package bootstrap

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tugas-sesi-10-arsitektur-berbasis-layanan/config"
	"tugas-sesi-10-arsitektur-berbasis-layanan/databases"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/app"
)

func Bootstrap() {
	db, err := databases.NewDatabase()
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	if err := db.SetupPool(); err != nil {
		log.Fatalf("DB pool setup failed: %v", err)
	}

	defer db.Close()

	logger := config.NewLogger()
	engine := config.NewEngine()
	api := engine.Group("api")
	v1 := app.App{
		DB:     db.DB,
		Engine: api,
		Log:    logger,
	}

	v1.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		port := fmt.Sprintf(":%s", "3291")
		if err = engine.Run(port); err != nil {
			logger.Error(err)
		}
	}()

	<-quit
	time.Sleep(1 * time.Second)

	log.Println("Shutting down server...")
}
