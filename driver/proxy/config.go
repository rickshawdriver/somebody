package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"net/http"
)

var (
	conf = &config.Config{}
)

func Start() {
	configure(conf)                                                           // configure config
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment // http agent

	p := NewProxy(conf)
	p.SetupSignal()
}

func configure(conf *config.Config) {
	if err := config.Load(conf); err != nil {
		log.Errorf("load config error: %s", err)
	}
	log.Initialize(&conf.Log)

	log.Info("conf init success")

	if err := config.WritePidToFile(); err != nil {
		log.Warnf("write pid file err %s", err)
	}
}
