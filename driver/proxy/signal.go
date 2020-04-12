package proxy

import (
	"github.com/rickshawdriver/somebody/pkg/log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

func (p *proxyRuntime) SetupSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	p.Stop()
	switch sig {
	case syscall.SIGUSR2: // restart
		log.Info("restart")
	case syscall.SIGINT, syscall.SIGTERM: // close
		signal.Stop(ch)
		close(ch)
		if err := p.FastHttpServer.Shutdown(); err != nil {
			log.Errorf("fastHttp close error is %s", err)
		}
		log.Info("fastHttp success close")
	}
}

func (p *proxyRuntime) Stop() {
	atomic.StoreInt32(&p.isStopped, 1)
}

func (p *proxyRuntime) IsStop() bool {
	return atomic.LoadInt32(&p.isStopped) == 1
}
