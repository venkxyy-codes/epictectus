package commands

import (
	"epictectus/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

var globalConfig = &config.Config{}

func SetupCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "to-do",
		Short:             "lb",
		PersistentPreRunE: commandSetup(),
		SilenceUsage:      true,
	}
	cmd.AddCommand(apiServerCommand())
	return cmd
}

func commandSetup() func(*cobra.Command, []string) error {
	return func(command *cobra.Command, strings []string) (err error) {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file, ", err)
			return err
		}
		return
	}
}
