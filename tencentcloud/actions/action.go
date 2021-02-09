package actions

import "fmt"

type Action interface {
	Service() string
	Action() string
	Version() string
	RateLimit()
	WithRateLimit() Action
	Host(domain string) string
}

type action struct {
	service   string
	action    string
	version   string
	rateLimit interface{}
}

func (act action) Service() string {
	return act.service
}

func (act action) Action() string {
	return act.action
}

func (act action) Version() string {
	return act.version
}

func (act action) RateLimit() {

}

func (act action) WithRateLimit() Action {
	return act
}

func (act action) Host(domain string) string {
	return fmt.Sprintf("%s.%s", act.service, domain)
}

var (
	CVMDescribeInstances = action{
		service:   "cvm",
		action:    "DescribeInstances",
		version:   "2017-03-12",
		rateLimit: nil,
	}
)
