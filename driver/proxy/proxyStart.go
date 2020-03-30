package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	conf config.Config
)

func Start(cmd *cobra.Command, args []string) {
	// init conf
	configure(&conf)
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment
}

func configure(conf *config.Config) {
	if err := config.Load(conf); err != nil {
		log.Errorf("load config error: %s")
	}
	// log init
	log.Initialize(&conf.Log)

	log.Infof("starting")
}
