package viper

import "github.com/juicyluv/online-store/internal/infrastructure/config"

func (c *configuration) LoggerLevel() string {
	return c.viper.GetString(config.LoggerLevel)
}

func (c *configuration) SetLoggerLevel(val string) {
	c.viper.Set(config.LoggerLevel, val)
}

func (c *configuration) LoggerLogToConsole() bool {
	return c.viper.GetBool(config.LoggerLogToConsole)
}

func (c *configuration) SetLoggerLogToConsole(val bool) {
	c.viper.Set(config.LoggerLogToConsole, val)
}

func (c *configuration) LoggerEncodeLogsAsJson() bool {
	return c.viper.GetBool(config.LoggerEncodeLogsAsJson)
}

func (c *configuration) SetLoggerEncodeLogsAsJson(val bool) {
	c.viper.Set(config.LoggerEncodeLogsAsJson, val)
}

func (c *configuration) LoggerLogToFile() bool {
	return c.viper.GetBool(config.LoggerLogToFile)
}

func (c *configuration) SetLoggerLogToFile(val bool) {
	c.viper.Set(config.LoggerLogToFile, val)
}

func (c *configuration) LoggerFilepath() string {
	return c.viper.GetString(config.LoggerFilepath)
}

func (c *configuration) SetLoggerFilepath(val string) {
	c.viper.Set(config.LoggerFilepath, val)
}

func (c *configuration) LoggerMaxLogFileSize() int {
	return c.viper.GetInt(config.LoggerMaxLogFileSize)
}

func (c *configuration) SetLoggerMaxLogFileSize(val int) {
	c.viper.Set(config.LoggerMaxLogFileSize, val)
}

func (c *configuration) LoggerMaxBackupsCount() int {
	return c.viper.GetInt(config.LoggerMaxBackupsCount)
}

func (c *configuration) SetLoggerMaxBackupsCount(val int) {
	c.viper.Set(config.LoggerMaxBackupsCount, val)
}

func (c *configuration) LoggerMaxLogFileAge() int {
	return c.viper.GetInt(config.LoggerMaxLogFileAge)
}

func (c *configuration) SetLoggerMaxLogFileAge(val int) {
	c.viper.Set(config.LoggerMaxLogFileAge, val)
}
