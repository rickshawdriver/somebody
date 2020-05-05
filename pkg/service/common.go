package service

import (
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/router"
	"github.com/rickshawdriver/somebody/store"
)

type Status uint32

const (
	Down Status = iota
	UP

	DEFAULTDEGREE = 3

	Page = int64(30)
)

type Dispatcher struct {
	Clusters map[uint32]*Cluster
	Services map[uint32]*Service
	Apis     map[uint32]*ApiRuntime

	Router *router.RootItem

	Store store.Store
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		Clusters: map[uint32]*Cluster{},
		Services: map[uint32]*Service{},
		Apis:     map[uint32]*ApiRuntime{},
		Router:   router.NewRouterList(DEFAULTDEGREE),
	}

	return d
}

func (d *Dispatcher) Load() {
	d.loadCluster()
	d.loadServices()
	d.loadApis()
}

func (d *Dispatcher) loadCluster() {
	log.Debug("load cluster ing .....")

	err := d.Store.Gets(Page, func() store.Pb {
		cluster := &Cluster{}
		return cluster
	}, func(value interface{}) error {
		if err := d.addCluster(value.(*Cluster)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err)
	}
}

func (d *Dispatcher) loadServices() {
	log.Debug("load service ing.....")

	err := d.Store.Gets(Page, func() store.Pb {
		service := &Service{}
		return service
	}, func(value interface{}) error {
		if err := d.AddService(value.(*Service)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err)
	}

	d.HealthCheckRun()
}

func (d *Dispatcher) loadApis() {
	log.Debug("load api ing.....")

	err := d.Store.Gets(Page, func() store.Pb {
		api := &ApiRuntime{}
		return api
	}, func(value interface{}) error {
		if err := d.addApi(value.(*ApiRuntime)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err)
	}
}
