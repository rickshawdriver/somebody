package cmd

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/spf13/cobra"
	"syscall"
)

func NewStopCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "stop",
		Short: "stop somebody",
		Long:  "stop somebody",
		Run: func(cmd *cobra.Command, args []string) {
			pid, err := config.Read()
			if err != nil {
				log.Error("get pid fail,please ensure startup")
			}
			if err := syscall.Kill(int(pid), syscall.SIGTERM); err != nil {
				log.Errorf("kill pid fail %s", err)
			}

			log.Info("exit somebody gateway success!!!")
		},
	}

	return versionCmd
}
