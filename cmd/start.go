package cmd

import (
	"github.com/spf13/cobra"
)

// start hordo
func Start() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start proxy",
		Long:  "startGateway",
		Run: func(cmd *cobra.Command, args []string) {
			//err := startProxy(cmd, args)
		},
	}

	return startCmd
}
