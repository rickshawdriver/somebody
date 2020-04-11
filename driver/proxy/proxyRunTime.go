package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/system"
	"github.com/rickshawdriver/somebody/router"
	"github.com/rickshawdriver/somebody/store"
)

type proxyRuntime struct {
	router    *router.RootItem
	dnsCache  *system.DnsCacheHandle
	Store     store.Store
	isStopped int32
}

func NewProxy(c *config.Config) *proxyRuntime {
	p := &proxyRuntime{
		router:   router.NewRouterList(c.RouterDegree),
		dnsCache: system.New(c.DnsCacheConf),
	}
	p.initStore()

	return p
}

func (p *proxyRuntime) initStore() {
	s, err := store.GetStoreFrom("etcd://127.0.0.1:2374", "hordo", "", "")
	log.Debug(err)
	if err != nil {
		log.Errorf("init store err %s", err)
	}

	p.Store = s
}
