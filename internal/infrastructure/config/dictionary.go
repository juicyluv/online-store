package config

const (
	KeyDelimiter = `.`

	LoggerLayer            = `logger`
	LoggerLevel            = LoggerLayer + KeyDelimiter + `level`
	LoggerLogToConsole     = LoggerLayer + KeyDelimiter + `log_to_console`
	LoggerEncodeLogsAsJson = LoggerLayer + KeyDelimiter + `encode_logs_as_json`
	LoggerLogToFile        = LoggerLayer + KeyDelimiter + `log_to_file`
	LoggerFilepath         = LoggerLayer + KeyDelimiter + `filepath`
	LoggerMaxLogFileSize   = LoggerLayer + KeyDelimiter + `max_log_file_size`
	LoggerMaxBackupsCount  = LoggerLayer + KeyDelimiter + `max_backups_count`
	LoggerMaxLogFileAge    = LoggerLayer + KeyDelimiter + `max_log_file_age`
)
