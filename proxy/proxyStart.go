package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/spf13/cobra"
)

var (
	startLogger = log.Get().WithField("prefix", "main")
)

func ServerStart(cmd *cobra.Command, args []string) {
	config.Global()
}
