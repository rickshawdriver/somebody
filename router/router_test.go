package router

import (
	"testing"
)

const (
	DEFAULTDEGREE = 3
)

func TestRouter(t *testing.T) {
	router := NewRouterList(DEFAULTDEGREE)

	router.Put([]byte("/test1"), 1001)
	router.Put([]byte("/test2"), 2001)
	router.Put([]byte("/test3"), 3001)
	router.Put([]byte("/test4"), 5001)
	router.Put([]byte("/test5"), 6001)
	router.Put([]byte("/test6"), 7001)

	id, found := router.Get([]byte("/test2"))
	if !found {
		t.Errorf("get router err, %t", found)
	}

	t.Logf("get router id, is %d", id)
}
