package system

import (
	"testing"
	"time"
)

func TestFetchItem(t *testing.T) {
	dc := DnsCacheConf{
		DnsRandomStrategy,
		time.Duration(10) * time.Second,
		time.Duration(5) * time.Second,
	}
	dnscache := New(dc)
	_, err := dnscache.FetchGet("www.lvchengchang.cn")
	if err != nil {
		t.Fatalf("dns parse got %s", err)
	}
}
