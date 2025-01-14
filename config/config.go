package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"ssl_mode"`
	} `mapstructure:"database"`
}

var Config AppConfig

func LoadConfig(filePath string) {
	viper.AutomaticEnv()                                   // Чтение переменных окружения
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Для замены . на _ в именах переменных

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Подставляем переменные окружения в конфигурацию
	for _, key := range viper.AllKeys() {
		val := viper.GetString(key)
		viper.Set(key, os.ExpandEnv(val)) // Подставляем переменные окружения
	}

	// Разбираем конфигурацию после подстановки переменных окружения
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
