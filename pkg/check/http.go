package check

import "time"

type HttpChecker struct {
	Name     string        `json:"name"`
	Endpoint string        `json:"endpoint"`
	Body     string        `json:"body"`
	Interval time.Duration `json:"interval"`
}

func Add(name, endpoint, body string, interval time.Duration) CheckerContainer {
	if Checkers == nil {
		Checkers = map[time.Duration][]Check{}
	}

	if Checkers[interval] == nil {
		Checkers[interval] = []Check{}
	}

	checker := HttpChecker{
		name, endpoint, body, interval,
	}

	Checkers[interval] = append(Checkers[interval], checker)
	return Checkers
}

func (h HttpChecker) Run() {

}
