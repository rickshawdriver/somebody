package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/store"
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
	config.SetGlobal(conf)

	if err := config.WritePidToFile(); err != nil {
		startLogger.Warnf("write PIDFile err:", err)
	}

	_, err := store.GetStoreFrom(config.GetStoreConf(conf.Store))
	if err != nil {
		startLogger.Warnf("err is:%s", err)
	}

	return nil
}
