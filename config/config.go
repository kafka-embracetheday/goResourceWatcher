package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"log"
	"os"
)

var (
	cfg    Config
	envMap = map[string]string{
		"release": "config",       // release 版本使用config.toml配置
		"debug":   "config.debug", // release 版本使用config.debug.toml配置
		"":        "config.debug", // 默认使用 debug 配置
	}
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	} `mapstructure:"app"`

	Logger struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
	} `mapstructure:"logger"`
}

func GetConfig() *Config {
	return &cfg
}

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "debug"
	}

	configName, exists := envMap[env]
	if !exists {
		log.Fatalf("Invalid APP_ENV value: %s", env)
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Printf("Unable to decode into struct: %v", err)
			return
		}
	})

}
