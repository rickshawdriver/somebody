package safe

import "github.com/rickshawdriver/somebody/pkg/log"

func Go(goroutine func()) {
	GoWithRecover(goroutine, log.LogPanicHandler)
}

func GoWithRecover(goroutine func(), panicHandle func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				panicHandle(err)
			}
		}()
		goroutine()
	}()
}
