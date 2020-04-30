package check

import (
	"github.com/rickshawdriver/somebody/pkg/safe"
	"time"
)

type Check interface {
	Run()
}

type CheckerContainer map[time.Duration][]Check

var (
	Checkers CheckerContainer

	SomebodyTimer []*time.Timer
)

func (cs CheckerContainer) RunCheck() {
	SomebodyTimer = []*time.Timer{}

	for k, _ := range cs {
		SomebodyTimer = append(SomebodyTimer, time.NewTimer(time.Second*k))
	}

	safe.Go(func() {

	})
}

func action(cc []Check, interval uint32) {

}
