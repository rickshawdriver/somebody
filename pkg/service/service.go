package service

import (
	"github.com/rickshawdriver/somebody/pkg/health"
	"github.com/rickshawdriver/somebody/pkg/log"
	"github.com/rickshawdriver/somebody/pkg/safe"
	"time"
)

type Service struct {
	ID          uint32              `json:"id"`
	EndPoint    string              `json:"endpoint"`
	MaxQps      uint32              `json:"max_qps"` // support max qps
	Status      Status              `json:"status"`
	HealthCheck *health.HttpChecker `json:"health_check"`
}

func (d *Dispatcher) AddService(service *Service) error {
	d.Services[service.ID] = service

	h := health.HttpChecker{
		service.ID,
		service.HealthCheck.Name,
		service.HealthCheck.Method,
		service.HealthCheck.Endpoint,
		service.HealthCheck.Body,
		service.HealthCheck.Interval,
	}
	health.AddCheck(&h)

	return nil
}

func (d *Dispatcher) HealthCheckRun() {
	if health.CheckContainer == nil {
		return
	}

	for interval, item := range health.CheckContainer {
		func(duration time.Duration, httpChecker []*health.HttpChecker, dispatcher *Dispatcher) {
			safe.Go(func() {
				doCheck(duration, httpChecker, dispatcher)
			})
		}(interval, item, d)
	}
}

func doCheck(interval time.Duration, h []*health.HttpChecker, d *Dispatcher) {
	if h == nil {
		return
	}

	s := time.NewTicker(interval * time.Second)
	for {
		select {
		case <-s.C:
			check(h, d)
		}
	}
}

func check(h []*health.HttpChecker, d *Dispatcher) {
	for _, checker := range h {
		if _, ok := d.Services[checker.ServiceId]; !ok {
			log.Warnf("dispatcher not have service: %d", checker.ServiceId)
			continue
		}
		status := UP
		if err := checker.Check(); err != nil {
			log.Warnf("health check error %s", err)
			status = Down
		}

		d.Services[checker.ServiceId].Status = status
	}
}

func (s *Service) GetID() uint32 {
	return s.ID
}
