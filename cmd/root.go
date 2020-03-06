package cmd

import (
	"github.com/spf13/cobra"
)

func GetRootCmd(args []string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "somebody",
		Short: "somebody gateway",
		Long:  `somebody gateway`,

		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	rootCmd.AddCommand(NewServerStartCmd())
	rootCmd.AddCommand(NewVersionCmd())
	rootCmd.AddCommand(NewCheckCmd())

	return rootCmd
}
