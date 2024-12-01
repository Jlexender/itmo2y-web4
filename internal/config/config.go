package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func InitConfig(path string) error {
	cfg = &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic(err)
	}
	return nil
}
