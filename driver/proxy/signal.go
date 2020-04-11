package proxy

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (p *proxyRuntime) SetupSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	switch sig {
	case syscall.SIGUSR2: // restart
		fmt.Println("restart")
	case syscall.SIGINT, syscall.SIGTERM: // close
		signal.Stop(ch)
		close(ch)
		fmt.Println("close")
	}
}
