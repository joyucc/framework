package config

import (
	"os"

	"github.com/micro/go-config"
	"github.com/micro/go-config/encoder/toml"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
)

var (
	configurationPath = "./config"
	currentConfig     = (*Config)(nil)
)

type Config struct {
	Database Database `json:"database"`
	Log      Log      `json:"log"`
	Http     HTTP     `json:"http"`
}

// SetConfigurationFile sets the path to the configuration file to use.
func SetConfigurationPath(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		panic("配置文件不存在")
	}
	configurationPath = path
}

// GetConfig returns the current configuration loaded.
func GetConfig() *Config {
	if currentConfig == nil {
		enc := toml.NewEncoder()

		// Load toml file with encoder
		if err := config.Load(file.NewSource(
			file.WithPath(configurationPath),
			source.WithEncoder(enc),
		)); err != nil {
			panic("加载Config文件出错...")
		}

		var (
			database Database
			log      Log
			http     HTTP
		)
		currentConfig = &Config{}
		if err := config.Get("database").Scan(database); err == nil {
			currentConfig.Database = database
		}
		if err := config.Get("log").Scan(log); err == nil {
			currentConfig.Log = log
		}
		if err := config.Get("http").Scan(http); err == nil {
			currentConfig.Http = http
		}

	}
	return currentConfig
}
