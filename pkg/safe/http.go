package safe

import (
	"github.com/prometheus/common/expfmt"
	"io"
	"net/http"
)

var (
	requestClient *http.Client
)

func Request(METHOD, pushUrl string, body io.Reader) (*http.Response, error) {
	if requestClient == nil {
		requestClient = http.DefaultClient
	}

	req, err := http.NewRequest(METHOD, pushUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", string(expfmt.FmtProtoDelim))
	resp, err := requestClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
