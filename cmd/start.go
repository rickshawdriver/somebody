package cmd

import (
	"github.com/rickshawdriver/somebody/driver/proxy"
	"github.com/spf13/cobra"
)

// start somebody
func NewServerStartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start proxy",
		Long:  "start Gateway",
		Run: func(cmd *cobra.Command, args []string) {
			proxy.Start(cmd, args)
		},
	}

	return startCmd
}
