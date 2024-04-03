package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env         string           `yaml:"env"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServerConfig `yaml:"http_server"`
	Database    DatabaseConfig   `yaml:"database"`
	Auth        AuthConfig       `yaml:"auth"`
}

type HttpServerConfig struct {
	Address     string `yaml:"address"`
	Timeout     int    `yaml:"timeout"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type DatabaseConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Name        string `yaml:"name"`
	MaxAttempts int    `yaml:"max_attempts"`
}

type AuthConfig struct {
	Secret string `yaml:"secret"`
}

func LoadConfig() *Config {
	const configPath = "config/local.yaml"

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Cannot read config, %s", err)

	}

	return &config
}
