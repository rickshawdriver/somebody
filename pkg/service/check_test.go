package service_test

import (
	"github.com/rickshawdriver/somebody/pkg/health"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/service"
	"testing"
	"time"
)

const (
	ENDPOINT  = "http://localhost:1080"
	ENDPOINT2 = "http://localhost:1081"

	METHOD = "POST"
)

func TestChecker(t *testing.T) {
	h1 := &health.HttpChecker{1, "hello,world", METHOD, ENDPOINT, "", 3}
	h2 := &health.HttpChecker{2, "hello,world", METHOD, ENDPOINT2, "", 6}
	service1 := &service.Service{
		1, "http://localhost:1080", 200, service.Down, h1,
	}
	service2 := &service.Service{
		2, "http://localhost:1081", 200, service.Down, h2,
	}

	d := service.NewDispatcher()
	d.AddService(service1)
	d.AddService(service2)
	d.HealthCheckRun()
	for {
		time.Sleep(time.Second * 1)
		log.Info(d.Services[1].Status)
		log.Info(d.Services[2].Status)
	}
}
