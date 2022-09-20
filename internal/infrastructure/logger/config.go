package logger

// loggerConfig is configuration for logging.
type loggerConfig struct {
	// LogToConsole enables console logging.
	LogToConsole bool

	// EncodeLogsAsJson makes the log framework use JSON format.
	EncodeLogsAsJson bool

	// LogToFile makes the framework log to a file.
	// The fields below can be skipped if this value is false.
	LogToFile bool

	// Filepath is the path of the log file.
	Filepath string

	// MaxLogFileSize is the max size in MB of the log file before it's rolled.
	MaxLogFileSize int

	// MaxBackupsCount is the max number of rolled files to keep.
	MaxBackupsCount int

	// MaxLogFileAge is the max age in days to keep a log file.
	MaxLogFileAge int
}
