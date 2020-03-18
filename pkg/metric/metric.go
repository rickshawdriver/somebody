package metric

import "time"

var (
	supportMetric = make(map[string]func(addr, namespace, instance string, interval time.Duration) (Metric, error))
)

type Metric interface {
	Request(string, string, time.Time)
	Report() error
	Statistics(int)
	Run()
}

func init() {
	supportMetric["prometheus"] = NewPrometheusMetric
}

func NewMetricInstance(metric, addr, namespace, instance string, interval time.Duration) (Metric, error) {
	metricInstance, ok := supportMetric[metric]
	if ok {
		return metricInstance(addr, namespace, instance, interval)
	}

	return nil, nil
}
