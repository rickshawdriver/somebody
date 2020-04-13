package store

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	supportSchema = make(map[string]func([]string, string, BasicAuth) (Store, error))
)

type BasicAuth struct {
	userName, passWord string
}

type StoreConf struct {
	RegistryAddr, NameSpace, UserName, Password string
}

type Store interface {
	Raw() interface{}

	Put(id uint32, f func() interface{}) error

	Get(id uint32, f func() interface{}) (interface{}, error)
	Gets(limit int64, types func() Pb, fn func(value interface{}) error) error
}

func init() {
	supportSchema["etcd"] = getEtcdStoreFrom
}

func GetStoreFrom(s StoreConf) (Store, error) {
	u, err := url.Parse(s.RegistryAddr)
	if err != nil {
		panic(fmt.Sprintf("parse registry addr failed, errors:%+v", err))
	}
	schema := strings.ToLower(u.Scheme)
	fn, ok := supportSchema[schema]
	if ok {
		return fn(getClusterAddr(u.Host), s.NameSpace, BasicAuth{userName: s.UserName, passWord: s.Password})
	}

	return nil, nil
}

func getClusterAddr(addr string) []string {
	var addrs []string
	values := strings.Split(addr, ",")

	for _, value := range values {
		addrs = append(addrs, fmt.Sprintf("http://%s", value))
	}

	return addrs
}

func getKey(namespace string, id uint32) string {
	return fmt.Sprintf("%s/%020d", namespace, id)
}
