package actions

type Action interface {
	Service() string
	Action() string
	RateLimit()
	WithRateLimit()
}

type action struct {
	service   string
	action    string
	rateLimit interface{}
}
