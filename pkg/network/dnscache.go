package network

import (
	"errors"
	cache "github.com/pmylund/go-cache"
	"math/rand"
	"net"
	"time"
)

const (
	DnsFirstStrategy  Strategy = "first"
	DnsRandomStrategy Strategy = "random"
	DnsNoCaheStrategy Strategy = "noCache"
)

type Strategy string

type DnsCacheHandle struct {
	dnsCacheStorage  *cache.Cache
	dnsCacheStrategy Strategy
	rand             *rand.Rand
}

func New(strategy Strategy, expiration, checkInterval time.Duration) *DnsCacheHandle {
	handle := &DnsCacheHandle{cache.New(expiration, checkInterval), strategy, nil}
	return handle
}

func (dnsCache *DnsCacheHandle) Get(host string) net.IP {
	item, ok := dnsCache.dnsCacheStorage.Get(host)
	if !ok {
		return nil
	}

	return item.(net.IP)
}

func (dns *DnsCacheHandle) Set(host string, ip net.IP) {
	dns.dnsCacheStorage.Set(host, ip, cache.DefaultExpiration)
}

func (dns *DnsCacheHandle) FetchGet(host string) (net.IP, error) {
	if host == "" {
		return nil, errors.New("host empty")
	}

	item, ok := dns.dnsCacheStorage.Get(host)
	if ok {
		return item.(net.IP), nil
	}

	ips, err := dns.resolveDNS(host)
	if err != nil {
		return nil, err
	}

	if dns.dnsCacheStrategy == DnsNoCaheStrategy {
		dns.dnsCacheStorage.Delete(host)
	}

	if dns.dnsCacheStrategy == DnsRandomStrategy {
		if len(ips) > 1 {
			return dns.getRandomIp(ips)
		}
	}

	return ips[0], nil
}

func (dns *DnsCacheHandle) getRandomIp(ips []net.IP) (net.IP, error) {
	if dns.dnsCacheStrategy != DnsRandomStrategy {
		return nil, errors.New("net random strategy")
	}

	if dns.rand == nil {
		source := rand.NewSource(time.Now().Unix())
		dns.rand = rand.New(source)
	}

	return ips[dns.rand.Intn(len(ips))], nil
}

func (dns *DnsCacheHandle) resolveDNS(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}

func (dns *DnsCacheHandle) Delete(host string) {
	dns.dnsCacheStorage.Delete(host)
}

func (dns *DnsCacheHandle) Clear() {
	dns.dnsCacheStorage.Flush()
}
