package config

import (
	"bitbucket.org/apps-for-bharat/gotools/blog"
	"os"
)

var appConfig *Config

type Config struct {
	AppName                string `yaml:"APP_NAME" env:"APP_NAME"`
	AppPort                string `yaml:"APP_PORT" env:"APP_PORT"`
	ENV                    string `yaml:"ENV" env:"ENVIRONMENT"`
	DbConfig               DbConfig
	RazorpayHttpConfig     RazorpayHttpConfig
	RazorpayCredentials    RazorpayCredentials
	LeadsquaredHttpConfig  LeadsquaredHttpConfig
	LeadsquaredCredentials LeadsquaredCredentials
	LogConfig              blog.LogConfig
}

func NewConfig() (*Config, error) {
	appConfig = &Config{}
	appConfig.SetDefault()
	return appConfig, nil
}

func (c *Config) SetDefault() {
	c.AppName = "epictectus"
	c.AppPort = "9090"
	c.ENV = "dev"
	c.DbConfig = DbConfig{
		Host:     "localhost",
		Username: "",
		Password: "",
		Port:     "27017",
		DBName:   "epictectus",
	}
	c.LogConfig = blog.LogConfig{
		Level:    "info",
		Format:   "json",
		Output:   "stdout",
		UnixTime: true,
	}
}

func GetConfig() *Config {
	return &Config{
		ENV:       os.Getenv("ENV"),
		DbConfig:  GetDbConfig(),
		AppName:   os.Getenv("APP_NAME"),
		AppPort:   os.Getenv("APP_PORT"),
		LogConfig: GetLogConfig(),
	}
}
