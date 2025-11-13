package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServerPort    int    `env:"HTTP_SERVER_PORT"`
	HTTPServerHost    string `env:"HTTP_SERVER_HOST"`
	PostgresUsername  string `env:"POSTGRES_USERNAME"`
	PostgresPassword  string `env:"POSTGRES_PASSWORD"`
	PostgresDB        string `env:"POSTGRES_DB"`
	PostgresPort      int    `env:"POSTGRES_PORT"`
	PostgresInnerHost string `env:"POSTGRES_HOST_INNER"`
	PostgresOuterHost string `env:"POSTGRES_HOST_OUTER"`
	SSLMode           string `env:"SSL_MODE"`
	ConfigPath        string `env:"CONFIG_PATH"`
}

// MustLoad - загрузка данных из .env в конфиг.
func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is not specified")
	}
	cfg := Config{}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(fmt.Sprintf("cannot read config from file; err=%s", err.Error()))
	}
	return &cfg
}

// fetchConfigPath - парсит пусть до файла с конфигом.
// Приоритет: значение из флага при запуске > дефолтное значение.
func fetchConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
	return configPath
}
