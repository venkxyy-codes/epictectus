package main

import (
	"context"
	"epictectus/commands"
	"os"
	"os/signal"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	cmd := commands.SetupCommands()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
