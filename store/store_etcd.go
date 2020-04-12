package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"math"
	"time"
)

const (
	DEFAULTTIMEOUT = 3 * time.Second

	endID = uint32(math.MaxUint32)
)

type EtcdStore struct {
	etcdRaw *clientv3.Client

	namespace string

	clusterDir string
}

type Pb interface {
	GetID() uint32
}

func getEtcdStoreFrom(addr []string, nameSpace string, auth BasicAuth) (Store, error) {
	etcd := &EtcdStore{
		namespace:  nameSpace,
		clusterDir: fmt.Sprintf("%s/clusters", nameSpace),
	}

	config := &clientv3.Config{
		Endpoints:   addr,
		DialTimeout: DEFAULTTIMEOUT,
	}

	setEtcdAuth(config, auth)

	client, err := clientv3.New(*config)

	if err != nil {
		return nil, err
	}

	etcd.etcdRaw = client
	return etcd, nil
}

func setEtcdAuth(config *clientv3.Config, auth BasicAuth) {
	if auth.userName != "" {
		config.Username = auth.userName
	}

	if auth.passWord != "" {
		config.Password = auth.passWord
	}
}

func (etcdStore *EtcdStore) Raw() interface{} {
	return etcdStore.etcdRaw
}

func (e *EtcdStore) get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(e.etcdRaw.Ctx(), DEFAULTTIMEOUT)
	defer cancel()

	return e.etcdRaw.Get(ctx, key, opts...)
}

func (e *EtcdStore) put(key, value string, opts ...clientv3.OpOption) error {
	_, err := e.etcdRaw.Put(e.etcdRaw.Ctx(), key, value, opts...)
	return err
}

func (e *EtcdStore) Put(id uint32, f func() interface{}) error {
	j, err := json.Marshal(f())
	if err != nil {
		return err
	}

	return e.put(getKey(e.namespace, id), string(j))
}

func (e *EtcdStore) Get(id uint32, f func() interface{}) (interface{}, error) {
	resp, err := e.get(getKey(e.namespace, id), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	value := f()
	if err = json.Unmarshal(resp.Kvs[0].Value, value); err != nil {
		return nil, err
	}

	return value, nil
}

func (e *EtcdStore) Gets(page int64, types func() Pb, fn func(value interface{}) error) error {
	start := uint32(0)
	end := getKey(e.namespace, endID)
	withRange := clientv3.WithRange(end)
	withLimit := clientv3.WithLimit(page)

	for {
		resp, err := e.get(getKey(e.namespace, start), withRange, withLimit)
		if err != nil {
			return err
		}

		for _, item := range resp.Kvs {
			value := types()
			if err = json.Unmarshal(item.Value, value); err != nil {
				return err
			}

			err = fn(value)
			if err != nil {
				return err
			}

			start = value.GetID() + 1
		}

		// read complete
		if len(resp.Kvs) < int(page) {
			break
		}
	}

	return nil
}
