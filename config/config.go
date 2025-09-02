package config

import (
	"os"
)

type Config struct {
	AppName  string   `yaml:"APP_NAME" env:"APP_NAME"`
	AppPort  string   `yaml:"APP_PORT" env:"APP_PORT"`
	ENV      string   `yaml:"ENV" env:"ENVIRONMENT"`
	DbConfig DbConfig `yaml:"DB_CONFIG" env:"DB_CONFIG"`
}

func (c *Config) SetDefault() {
	c.AppName = "to-do"
	c.AppPort = "9090"
	c.ENV = "dev"
	c.DbConfig = DbConfig{
		Host:     "localhost",
		Username: "",
		Password: "",
		Port:     "27017",
		DBName:   "to-do",
	}
}

func GetConfig() *Config {
	return &Config{
		ENV:      os.Getenv("ENV"),
		DbConfig: GetDbConfig(),
		AppName:  os.Getenv("APP_NAME"),
		AppPort:  os.Getenv("APP_PORT"),
	}
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
