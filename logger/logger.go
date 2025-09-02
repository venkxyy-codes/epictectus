package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	//can add for development as well
	//logger, err = zap.NewProduction()
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}
