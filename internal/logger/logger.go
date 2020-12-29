package logger

import (
	"crypto/tls"
	"github.com/rs/zerolog"
	"github.com/zerodays/sistem-payments/internal/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Shared logger instance.
var Log zerolog.Logger

// Init initializes the shared logger instance.
func Init() {
	// Create writers that are writing logs.
	var writers []io.Writer

	if config.Logs.ConsoleLogging() {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.Logs.FileLogging() {
		writers = append(writers, newRollingFile())
	}
	if config.Logs.LogitLogging() {
		conn, err := tls.Dial("tcp", config.Logs.LoggitURL(), &tls.Config{RootCAs: nil})
		if err != nil {
			log.Fatal(err)
		}

		writers = append(writers, conn)
	}

	multiWriter := io.MultiWriter(writers...)

	// Configure shared logger instance.
	Log = zerolog.New(multiWriter).Level(zerolog.Level(config.Logs.LogLevel())).
		With().
		Str("app_name", "sistem_payments").
		Str("version", "0.0.1").
		Timestamp().
		Caller().
		Logger()
}

// newRollingFile creates io.Writer object that writes into photos
// and rolls it over after size or age exceeds the specified values.
func newRollingFile() io.Writer {
	// Create logs directory.
	directory := filepath.Dir(config.Logs.LogPath())
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Create new rolling photos writer.
	return &lumberjack.Logger{
		Filename:   config.Logs.LogPath(),
		MaxSize:    config.Logs.MaxSize(),
		MaxAge:     config.Logs.MaxAge(),
		MaxBackups: config.Logs.MaxBackups(),
	}
}
