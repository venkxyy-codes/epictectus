package main

import (
	"context"
	"epictectus/commands"
	"epictectus/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	logger.InitLogger()
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger.Logger)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	cmd := commands.SetupCommands()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
