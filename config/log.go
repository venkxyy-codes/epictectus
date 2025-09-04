package config

import (
	"bitbucket.org/apps-for-bharat/gotools/blog"
	"os"
)

func GetLogConfig() blog.LogConfig {
	return blog.LogConfig{
		Level:    os.Getenv("LOG_LEVEL"),
		Format:   os.Getenv("LOG_FORMAT"),
		Output:   os.Getenv("LOG_OUTPUT"),
		UnixTime: false,
	}
}
