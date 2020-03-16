package metric

import "time"

var (
	supportMetric = make(map[string]func(addr, namespace, instance string) (Metric, error))
)

type Metric interface {
	Request(string, string, time.Time)
	Report() error
	Statistics(int)
}

func init() {
	supportMetric["prometheus"] = NewPrometheusMetric
}

func NewMetricInstance(metric, addr, namespace, instance string) (Metric, error) {
	metricInstance, ok := supportMetric[metric]
	if ok {
		return metricInstance(addr, namespace, instance)
	}

	return nil, nil
}
