package commands

import (
	"epictectus/appcontext"
	"epictectus/config"
	"epictectus/logger"
	"epictectus/router"
	"epictectus/service"
	"fmt"
	"github.com/spf13/cobra"
)

func apiServerCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "start_http",
		Short: "Run the determine web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Logger.Info("initialising server command")
			globalConfig := config.GetConfig()

			appcontext.Init()
			db := appcontext.GetDBClient()
			fmt.Println(db)

			serverDependencies := service.InstantiateServerDependencies()
			r := router.InitRouter(router.Options{
				Conf:         globalConfig,
				Dependencies: serverDependencies,
			})

			port := fmt.Sprintf(":%s", globalConfig.AppPort)
			fmt.Println("Listening on port: ", port, globalConfig.AppPort)
			err := r.Run(port)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return command
}
