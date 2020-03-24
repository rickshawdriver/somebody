package cmd

import (
	"github.com/rickshawdriver/somebody/driver/api"
	"github.com/spf13/cobra"
)

// start api server
func NewApiServerStartCmd() *cobra.Command {
	apiServerCmd := &cobra.Command{
		Use:   "api",
		Short: "api server start",
		Long:  "start api server",
		Run: func(cmd *cobra.Command, args []string) {
			api.Start(cmd, args)
		},
	}

	return apiServerCmd
}
