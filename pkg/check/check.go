package check

import (
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/safe"
	"time"
)

type CheckerContainer map[time.Duration][]HttpChecker

var (
	Checkers CheckerContainer
)

func (cs CheckerContainer) RunCheck() {
	for _, item := range cs {
		safe.Go(func() {
			t := time.NewTicker(time.Second * 5)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					for _, value := range item {
						err := value.Run()
						if err != nil {
							log.Warnf("health check error %s", err)
						}
					}
				}
			}
		})
	}

	for {

	}
}
