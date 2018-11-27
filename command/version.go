package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const goottServerVersion = 0.1

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "list the version of goott-server",
		Long:  `list the version of goott-server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(goottServerVersion)
		},
	}

	return cmd
}
