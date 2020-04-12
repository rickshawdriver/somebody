package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

const (
	DEFAULTTIMEOUT = 3 * time.Second
)

type EtcdStore struct {
	etcdRaw *clientv3.Client

	namespace string

	clusterDir string
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
