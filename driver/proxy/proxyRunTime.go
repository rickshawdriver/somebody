package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/safe"
	"github.com/rickshawdriver/somebody/pkg/service"
	"github.com/rickshawdriver/somebody/pkg/system"
	"github.com/rickshawdriver/somebody/router"
	"github.com/rickshawdriver/somebody/store"
	"github.com/valyala/fasthttp"
)

type proxyRuntime struct {
	router   *router.RootItem
	dnsCache *system.DnsCacheHandle
	Store    store.Store
	Conf     *config.Config

	FastHttpServer *fasthttp.Server

	isStopped int32
}

func NewProxy(c *config.Config) *proxyRuntime {
	p := &proxyRuntime{
		router:   router.NewRouterList(c.RouterDegree),
		dnsCache: system.New(c.DnsCacheConf),
		Conf:     c,
	}

	if d := p.initStore().load().NewHttpServer(); d == nil {
		log.Error("init proxy error")
	}

	return p
}

func (p *proxyRuntime) initStore() *proxyRuntime {
	s, err := store.GetStoreFrom(p.Conf.Store)

	if err != nil {
		log.Errorf("init store err %s", err)
	}

	p.Store = s
	return p
}

func (p *proxyRuntime) load() *proxyRuntime {
	p.loadCluster()

	return p
}

func (p *proxyRuntime) NewHttpServer() *proxyRuntime {
	p.FastHttpServer = &fasthttp.Server{
		Handler: p.HttpServerHandle,
	}

	safe.Go(func() {
		if err := p.FastHttpServer.ListenAndServe(p.Conf.HttpConf.Addr); err != nil {
			log.Errorf("listen fastHttp err %s", err)
		}
	})

	return p
}

func (p *proxyRuntime) HttpServerHandle(ctx *fasthttp.RequestCtx) {
	if p.IsStop() {
		log.Warn("fastHttp already stop")
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return
	}
	log.Debug(string(ctx.Request.RequestURI()))
}

func (p *proxyRuntime) loadCluster() {
	log.Debug("load clustering .....")

	data, err := p.Store.Get(1, func() interface{} {
		cluster := &service.Cluster{}
		return cluster
	})

	if err != nil {
		log.Error(err)
	}

	log.Info(data.(*service.Cluster).Name)
	log.Info(int64(32))
}
