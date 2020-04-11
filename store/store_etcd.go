package store

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

const (
	DEFAULTTIMEOUT = 3
)

type EtcdStore struct {
	etcdRaw *clientv3.Client

	namespace string
}

func getEtcdStoreFrom(addr []string, nameSpace string, auth BasicAuth) (Store, error) {
	etcd := &EtcdStore{
		namespace: nameSpace,
	}

	config := &clientv3.Config{
		Endpoints:   addr,
		DialTimeout: DEFAULTTIMEOUT * time.Second,
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

func (etcdStore *EtcdStore) Ping() error {
	return nil
}
