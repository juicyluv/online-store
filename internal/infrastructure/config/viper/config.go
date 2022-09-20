package viper

import (
	"github.com/juicyluv/online-store/internal/infrastructure/config"
	"github.com/spf13/viper"
	"path/filepath"
)

type configuration struct {
	viper *viper.Viper
}

func NewConfig(configFullFilePath string) (config.Config, error) {
	cfg := &configuration{
		viper: viper.NewWithOptions(viper.KeyDelimiter(config.KeyDelimiter)),
	}

	info := getConfigFileInfoFromPath(configFullFilePath)

	cfg.viper.AddConfigPath(info.FilePath)
	cfg.viper.SetConfigName(info.FileName)
	cfg.viper.SetConfigType(info.FileExt)

	err := cfg.viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type configFileInfo struct {
	FileName string
	FileExt  string
	FilePath string
}

func getConfigFileInfoFromPath(path string) configFileInfo {
	if path == "" {
		return configFileInfo{}
	}

	dir, filename := filepath.Split(path)

	return configFileInfo{
		FileName: filename,
		FileExt:  filepath.Ext(filename),
		FilePath: dir,
	}
}
