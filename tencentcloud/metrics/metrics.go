package metrics

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

var (
	RequestTotal = prom.NewCounterVec(
		prom.CounterOpts{Subsystem: "awesome_tencentcloud_client", Name: "request_total"},
		[]string{"sid", "service", "action", "version"},
	)
	ErrorTotal = prom.NewCounterVec(
		prom.CounterOpts{Subsystem: "awesome_tencentcloud_client", Name: "error_total"},
		[]string{"sid", "service", "action", "version", "code"},
	)
	NetworkFailureRetryTotal = prom.NewCounterVec(
		prom.CounterOpts{Subsystem: "awesome_tencentcloud_client", Name: "network_failure_retry_total"},
		[]string{"service", "action", "version"},
	)
	RateLimitRetryTotal = prom.NewCounterVec(
		prom.CounterOpts{Subsystem: "awesome_tencentcloud_client", Name: "rate_limit_retry_total"},
		[]string{"service", "action", "version"},
	)
	RequestDuration = prom.NewGaugeVec(
		prom.GaugeOpts{Subsystem: "awesome_tencentcloud_client", Name: "request_duration"},
		[]string{"service", "action", "version"},
	)
)
