package config

import "os"

type AngoorHttpConfig struct {
	Host           string `env:"ANGOOR_HOST"`
	TriggerWebhook string `env:"ANGOOR_TRIGGER_WEBHOOK"`
}

type AngoorCredentials struct {
	AccessKey string `env:"ANGOOR_API_KEY"`
}

func GetAngoorCredentials() AngoorCredentials {
	return AngoorCredentials{
		AccessKey: os.Getenv("ANGOOR_API_KEY"),
	}
}

func GetAngoorHttpConfig() AngoorHttpConfig {
	return AngoorHttpConfig{
		Host:           os.Getenv("ANGOOR_HOST"),
		TriggerWebhook: os.Getenv("ANGOOR_TRIGGER_WEBHOOK"),
	}
}
