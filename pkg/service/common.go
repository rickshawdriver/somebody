package service

import (
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/store"
)

type Status uint32

const (
	UP Status = iota
	Down

	ClusterPage = int64(30)
)

type Dispatcher struct {
	Clusters map[uint32]*Cluster

	Store store.Store
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		Clusters: map[uint32]*Cluster{},
	}

	return d
}

func (d *Dispatcher) Load() {
	d.loadCluster()
}

func (d *Dispatcher) loadCluster() {
	log.Debug("load clustering .....")

	err := d.Store.Gets(ClusterPage, func() store.Pb {
		cluster := &Cluster{}
		return cluster
	}, func(value interface{}) error {
		if err := d.AddCluster(value.(*Cluster)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err)
	}
}
