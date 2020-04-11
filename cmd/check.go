package cmd

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/spf13/cobra"
)

// start somebody
func NewCheckCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "check somebody of configuration file",
		Long:  "check config",
		Run: func(cmd *cobra.Command, args []string) {
			err := check()
			if err != nil {
				log.Errorf("configuration file not valid, err: %s", err)
			}
			log.Info("configuration file is valid")
		},
	}

	return checkCmd
}

func check() error {
	conf := &config.Config{}
	err := config.Load(conf)
	return err
}
