package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/config"
	"sync"
)

type proxyRuntime struct {
	sync.Mutex

	conf func() config.Config
}
