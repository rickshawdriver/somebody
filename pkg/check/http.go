package check

import (
	"bytes"
	"fmt"
	"github.com/rickshawdriver/somebody/pkg/safe"
	"net/http"
	"time"
)

var (
	CheckerClient *http.Client
)

type HttpChecker struct {
	Name     string        `json:"name"`
	Method   string        `json:"method"`
	Endpoint string        `json:"endpoint"`
	Body     string        `json:"body"`
	Interval time.Duration `json:"interval"`
}

func Add(name, endpoint, body, method string, interval time.Duration) CheckerContainer {
	if Checkers == nil {
		Checkers = map[time.Duration][]HttpChecker{}
	}

	if Checkers[interval] == nil {
		Checkers[interval] = []HttpChecker{}
	}

	checker := HttpChecker{
		name, method, endpoint, body, interval,
	}

	Checkers[interval] = append(Checkers[interval], checker)
	return Checkers
}

func (h HttpChecker) Run() error {
	if CheckerClient == nil {
		CheckerClient = http.DefaultClient
	}

	r := bytes.NewReader([]byte(h.Body))
	resp, err := safe.Request(h.Method, h.Endpoint, r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("health check status code not 200 %d", resp.StatusCode)
	}

	return nil
}
