package api

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/service"
	"github.com/rickshawdriver/somebody/pkg/system"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	apiLogger = log.Get()
)

func Start(cmd *cobra.Command, args []string) {
	var conf config.Config
	if err := config.GetConf(&conf); err != nil {
		apiLogger.Errorf("get conf err: %s", err)
	}

	router := service.NewHTTPServer()
	http.ListenAndServe(":8080", router)
	system.SetupSignal()
}
