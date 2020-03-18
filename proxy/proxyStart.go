package proxy

import (
	"fmt"
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/metric"
	"github.com/rickshawdriver/somebody/store"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	startLogger = log.Get().WithField("prefix", "main")
)

func ServerStart(cmd *cobra.Command, args []string) {
	if err := initialize(); err != nil {
		startLogger.Errorf("initialize error for %s", err)
	}

	setupSignal()
}

func initialize() error {
	var conf config.Config
	if err := config.Load(&conf); err != nil {
		return err
	}
	config.SetGlobal(conf)

	if err := config.WritePidToFile(); err != nil {
		startLogger.Warnf("write PIDFile err: %s", err)
	}

	_, err := store.GetStoreFrom(config.GetStoreConf(conf.Store))
	if err != nil {
		startLogger.Errorf("err is:%s", err)
	}

	p, err := metric.NewMetricInstance(conf.Metric.Type, conf.Metric.Addr, conf.Metric.Namespace,
		conf.Metric.Instance, conf.Metric.Interval)
	if err != nil {
		startLogger.Errorln(err)
	}
	p.Run()

	return nil
}

func setupSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	switch sig {
	case syscall.SIGUSR2: // restart
		fmt.Println("restart")
	case syscall.SIGINT, syscall.SIGTERM: // close
		signal.Stop(ch)
		close(ch)
		fmt.Println("close")
	}
}
