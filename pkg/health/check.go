package health

import (
	"bytes"
	"fmt"
	"github.com/rickshawdriver/somebody/pkg/safe"
	"net/http"
	"time"
)

var (
	CheckContainer map[time.Duration][]*HttpChecker

	CheckRequestClient *http.Client
)

type HttpChecker struct {
	ServiceId uint32        `json:"service_id"`
	Name      string        `json:"name"`
	Method    string        `json:"method"`
	Endpoint  string        `json:"endpoint"`
	Body      string        `json:"body"`
	Interval  time.Duration `json:"interval"`
}

func AddCheck(h *HttpChecker) {
	httpChecker := h

	if CheckContainer == nil {
		CheckContainer = map[time.Duration][]*HttpChecker{}
	}

	CheckContainer[h.Interval] = append(CheckContainer[h.Interval], httpChecker)
}

func (h *HttpChecker) Check() error {
	if CheckRequestClient == nil {
		CheckRequestClient = http.DefaultClient
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
