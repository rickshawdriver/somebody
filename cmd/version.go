package cmd

import (
	"fmt"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/spf13/cobra"
	"runtime"
)

const BINARY = "0.0.1"

func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, BINARY, runtime.Version())
}

func NewVersionCmd() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print Version",
		Long:  "Print Now Gateway Version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Get().Infoln(String("somebody"))
		},
	}

	return versionCmd
}
