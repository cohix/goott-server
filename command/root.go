package command

import (
	"fmt"
	"os"

	"github.com/cohix/goott-server/model"
	"github.com/cohix/goott-server/service"
	log "github.com/cohix/simplog"
	"github.com/spf13/cobra"
)

// Execute runs the tool
func Execute() {
	cmd := arrangeCommands()

	if err := cmd.Execute(); err != nil {
		log.LogError(err)
		os.Exit(1)
	}
}

func arrangeCommands() *cobra.Command {
	root := rootCmd()
	root.AddCommand(versionCmd())

	return root
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goott-server",
		Short: "goott-server is an example command-line client for a gRPC service",
		Long:  `An example of building command-line clients in Go, and a primer for gRPC`,
		Run: func(cmd *cobra.Command, args []string) {
			errChan := make(chan error)

			app := model.InitializeApp()

			log.LogInfo("starting goott server on :3687")
			log.LogInfo(fmt.Sprintf("to authenticate goott: `export GOOTT_TOKEN=%s", app.AuthToken))
			go service.StartGoottService(app, errChan)

			for {
				if err := <-errChan; err != nil {
					log.LogError(err)
					os.Exit(1)
				}
			}
		},
	}

	return cmd
}
