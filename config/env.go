package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Mode  string
	DBDSN string
)

func init() {
	_ = godotenv.Load()
	Mode = os.Getenv("Mode")
	DBDSN = os.Getenv("DB_DSN")
}
