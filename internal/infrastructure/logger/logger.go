package logger

import (
	"github.com/juicyluv/online-store/internal/infrastructure/config"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

// Logger represents zerolog.Logger interface.
type Logger struct {
	*zerolog.Logger
}

// New returns a new logger instance.
func New(config config.Config) *Logger {
	cfg := &loggerConfig{
		LogToConsole:     config.LoggerLogToConsole(),
		EncodeLogsAsJson: config.LoggerEncodeLogsAsJson(),
		LogToFile:        config.LoggerLogToFile(),
		Filepath:         config.LoggerFilepath(),
		MaxLogFileSize:   config.LoggerMaxLogFileSize(),
		MaxBackupsCount:  config.LoggerMaxBackupsCount(),
		MaxLogFileAge:    config.LoggerMaxLogFileAge(),
	}

	var writers []io.Writer

	if cfg.LogToConsole {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.LogToFile {
		writers = append(writers, newRollingFile(cfg))
	}

	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("log_to_file", cfg.LogToFile).
		Bool("json_log", cfg.EncodeLogsAsJson).
		Str("log_filepath", cfg.Filepath).
		Int("max_backups", cfg.MaxBackupsCount).
		Int("max_age_days", cfg.MaxLogFileAge).
		Msg("Logger has been configured.")

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(config *loggerConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Filepath,
		MaxBackups: config.MaxBackupsCount,
		MaxSize:    config.MaxLogFileSize,
		MaxAge:     config.MaxLogFileAge,
	}
}
