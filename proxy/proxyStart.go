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
	if err := initialize(); err != nil {
		startLogger.Errorf("initialize error for %s", err)
	}

}

func initialize() error {
	var conf config.Config
	if err := config.Load(&conf); err != nil {
		return err
	}

	if err := config.WritePidToFile(); err != nil {
		startLogger.Warnf("write PIDFile err:", err)
	}

	return nil
}
