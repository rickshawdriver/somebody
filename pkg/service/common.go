package service

import (
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/store"
)

type Status uint32

const (
	Down Status = iota
	UP

	Page = int64(30)
)

type Dispatcher struct {
	Clusters map[uint32]*Cluster
	Services map[uint32]*Service

	Store store.Store
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		Clusters: map[uint32]*Cluster{},
		Services: map[uint32]*Service{},
	}

	return d
}

func (d *Dispatcher) Load() {
	d.loadCluster()
	d.loadServices()
}

func (d *Dispatcher) loadCluster() {
	log.Debug("load clustering .....")

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
	log.Debug("load serviceing.....")

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
