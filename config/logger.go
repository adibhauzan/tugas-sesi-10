package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var hari = []string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu",
}

var bulan = []string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

type CustomFormatter struct {
	ForceColors  bool
	PadLevelText bool
	UseIndonesia bool
}

func prettyJSONIfPossible(s string) string {
	var buf bytes.Buffer

	if json.Valid([]byte(s)) {
		err := json.Indent(&buf, []byte(s), "", "  ")

		if err == nil {
			return buf.String()
		}
	}

	return s
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	t := entry.Time.Local()

	var waktu string

	if f.UseIndonesia {
		waktu = fmt.Sprintf("%s, %02d %s %d %02d:%02d:%02d",
			hari[t.Weekday()],
			t.Day(),
			bulan[int(t.Month())-1],
			t.Year(),
			t.Hour(), t.Minute(), t.Second(),
		)
	} else {
		waktu = t.Format(time.RFC3339)
	}

	level := entry.Level.String()

	if f.PadLevelText {
		level = fmt.Sprintf("%-6s", level)
	}

	msg := fmt.Sprintf("%s [%s] %s\n", level, waktu, entry.Message)

	for k, v := range entry.Data {

		if str, ok := v.(string); ok {
			msg += fmt.Sprintf("    %s: %v\n", k, prettyJSONIfPossible(str))
		} else {
			msg += fmt.Sprintf("    %s: %v\n", k, v)
		}
	}

	return []byte(msg), nil
}

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	var level logrus.Level
	switch Mode {
	case "local":
		level = logrus.DebugLevel
	case "production":
		level = logrus.InfoLevel
	case "staging":
		level = logrus.DebugLevel
	case "development":
		level = logrus.DebugLevel
	default:
		level = logrus.InfoLevel
	}
	log.SetFormatter(&CustomFormatter{
		ForceColors:  true,
		PadLevelText: true,
		UseIndonesia: true,
	})

	log.SetLevel(level)
	return log
}
