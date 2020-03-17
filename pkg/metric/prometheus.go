package metric

import (
	"bytes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

const (
	METHOD = "PUT"
)

type Prometheus struct {
	job, instance, addr string
	requestClient       *http.Client
	interval            time.Duration
}

var (
	requestContainer = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "somebody",
		Subsystem: "somebody",
		Name:      "somebody_request",
		Help:      "gateway somebody api",
	})

	requestCost = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "somebody",
			Subsystem: "somebody",
			Name:      "somebody_duration_seconds",
			Help:      "the cost of somebody",
		},
		[]string{"apiname"},
	)

	qpsTarget = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "somebody",
			Help: "moment",
		},
		[]string{"server"},
	)
)

func init() {
	prometheus.MustRegister(requestContainer, requestCost, qpsTarget)
}

func NewPrometheusMetric(addr, namespace, instance string, interval time.Duration) (Metric, error) {
	client := http.DefaultClient
	client.Timeout = time.Second * 3

	p := &Prometheus{
		requestClient: client,
		job:           namespace,
		instance:      instance,
		addr:          addr,
		interval:      interval,
	}

	return p, nil
}

func (p *Prometheus) Request(api, code string, startTime time.Time) {
	requestContainer.Inc()
	now := time.Now()
	requestCost.WithLabelValues(api).Observe(now.Sub(startTime).Seconds())
}

func (p *Prometheus) Statistics(qps int) {
	qpsTarget.WithLabelValues("qps").Set(float64(qps))
}

func (p *Prometheus) Report() error {
	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	enc := expfmt.NewEncoder(buf, expfmt.FmtProtoDelim)

	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(METHOD, p.addr, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", string(expfmt.FmtProtoDelim))
	resp, err := p.requestClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (p *Prometheus) Run() {
	go func() {
		t := time.NewTicker(p.interval * time.Second)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				err := p.Report()
				// todo wait fix push url error
				if err != nil {
					log.Errorf("metric: could not push metrics to prometheus pushgateway: errors:\n%+v", err)
				}
			}
		}
	}()
}
