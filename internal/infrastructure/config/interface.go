package config

type Config interface {
	logger
}

type logger interface {
	LoggerLevel() string
	SetLoggerLevel(val string)

	LoggerLogToConsole() bool
	SetLoggerLogToConsole(val bool)

	LoggerEncodeLogsAsJson() bool
	SetLoggerEncodeLogsAsJson(val bool)

	LoggerLogToFile() bool
	SetLoggerLogToFile(val bool)

	LoggerFilepath() string
	SetLoggerFilepath(val string)

	LoggerMaxLogFileSize() int
	SetLoggerMaxLogFileSize(val int)

	LoggerMaxBackupsCount() int
	SetLoggerMaxBackupsCount(val int)

	LoggerMaxLogFileAge() int
	SetLoggerMaxLogFileAge(val int)
}
