package network

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchItem(t *testing.T) {
	dnscache := New(DnsRandomStrategy, time.Duration(10)*time.Second, time.Duration(5)*time.Second)
	ip, err := dnscache.FetchGet("www.lvchengchang.cn")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip)
}
