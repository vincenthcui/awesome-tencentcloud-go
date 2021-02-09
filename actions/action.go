package actions

type Action interface {
	Service() string
	Action() string
	Version() string
	RateLimit()
	WithRateLimit()
	Host(domain string) string
}

type action struct {
	service   string
	action    string
	rateLimit interface{}
}
