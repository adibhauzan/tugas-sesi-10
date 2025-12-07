package databases

import (
	"fmt"
	"log"
	"os"
	"time"
	"tugas-sesi-10-arsitektur-berbasis-layanan/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {

	var level logger.LogLevel
	switch config.Mode {
	case "production":
		level = logger.Error
	case "staging":
		level = logger.Warn
	case "development":
		level = logger.Info
	case "local":
		level = logger.Info
	default:
		level = logger.Info
	}

	db, err := gorm.Open(postgres.Open(config.DBDSN), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  level,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		FullSaveAssociations: true,
		PrepareStmt:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *Database) SetupPool() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return nil
}
