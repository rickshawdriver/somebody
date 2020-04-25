package check

type HttpChecker struct {
	name     string `json:"name"`
	endpoint string `json:"endpoint"`
	body     string `json:"body"`
	interval int    `json:"interval"`
}
