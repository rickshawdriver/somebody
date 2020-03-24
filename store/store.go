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

type Store interface {
	Ping() error
}

func init() {
	supportSchema["etcd"] = getEtcdStoreFrom
}

func GetStoreFrom(registryAddr, nameSpace, userName, password string) (Store, error) {
	u, err := url.Parse(registryAddr)
	if err != nil {
		panic(fmt.Sprintf("parse registry addr failed, errors:%+v", err))
	}
	schema := strings.ToLower(u.Scheme)
	fn, ok := supportSchema[schema]
	if ok {
		return fn(getClusterAddr(u.Host), nameSpace, BasicAuth{userName: userName, passWord: password})
	}

	return nil, nil
}

func getClusterAddr(addr string) []string {
	var addrs []string
	values := strings.Split(addr, ",")

	for _, value := range values {
		addrs = append(addrs, fmt.Sprintf("service://%s", value))
	}

	return addrs
}
